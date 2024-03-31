package logminder

import (
	"time"
)

type User struct {
	ID       string                 `json:"id"`
	FullName string                 `json:"full_name"`
	Role     string                 `json:"role"`
	Details  map[string]interface{} `json:"details"`
}

type Action struct {
	Type    string                 `json:"type"`
	Model   string                 `json:"model"`
	Details map[string]interface{} `json:"details"`
}

type WriteLogPayload struct {
	User    User   `json:"user"`
	Action  Action `json:"action"`
	Message string `json:"message"`
}

type RetreiveLogsFilter struct {
	Limit  int32                  `json:"limit" form:"limit"`
	Offset int32                  `json:"offset" form:"offset"`
	Query  map[string]interface{} `json:"query" form:"query"`
}

type Log struct {
	ID        string    `json:"id"`
	User      User      `json:"user"`
	Action    Action    `json:"action"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type RetreiveLogsResponse struct {
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
	Total      int64 `json:"total"`
	TotalPages int64 `json:"total_pages"`
	Logs       []Log `json:"logs"`
}
