package logminder

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type httplogminder struct {
	client *resty.Client
	cfg    *config
}

func NewHttpClient(cfg *config) LogMinderClient {
	return &httplogminder{
		client: resty.New().
			SetDebug(cfg.debug).
			SetJSONMarshaler(json.Marshal).
			SetJSONUnmarshaler(json.Unmarshal).
			SetBaseURL(cfg.addr).
			SetBasicAuth(cfg.username, cfg.password),
		cfg: cfg,
	}
}

func (l *httplogminder) WriteLog(ctx context.Context, payload *WriteLogPayload) error {

	if payload.User.ID == "" {
		return ErrUserIDInvalid
	}

	if payload.User.Role == "" {
		return ErrUserRoleInvalid
	}

	if payload.Action.Type == "" {
		return ErrActionTypeInvalid
	}

	go func() {
		// write logs
		l.client.R().SetBody(payload).Post("v1/logs")
	}()

	return nil
}

func (l *httplogminder) RetreiveLogs(ctx context.Context, payload *RetreiveLogsFilter) (*RetreiveLogsResponse, error) {

	var resp RetreiveLogsResponse

	res, err := l.client.R().SetContext(ctx).SetResult(&resp).SetQueryParams(map[string]string{
		"limit":          fmt.Sprintf("%d", payload.Limit),
		"offset":         fmt.Sprintf("%d", payload.Offset),
		"search":         payload.Search,
		"user_id":        payload.UserID,
		"user_role":      payload.UserRole,
		"from":           payload.FromDate.Format("2006-01-02"),
		"to":             payload.ToDate.Format("2006-01-02"),
		"action_type":    payload.ActionType,
		"action_details": payload.ActionDetails,
		"user_details":   payload.UserDetails,
	}).Get("/v1/logs")

	if err != nil {
		return nil, ErrCannotMakeRequest
	}

	if res.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("log-minder response code: %d, and body: %s", res.StatusCode(), string(res.Body()))
	}

	return &resp, nil
}
