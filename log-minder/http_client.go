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

	res, err := l.client.R().SetContext(ctx).SetResult(&resp).SetBody(payload).Post("/v1/logs/retreive")

	if err != nil {
		return nil, ErrCannotMakeRequest
	}

	if res.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("log-minder response code: %d, and body: %s", res.StatusCode(), string(res.Body()))
	}

	return &resp, nil
}
