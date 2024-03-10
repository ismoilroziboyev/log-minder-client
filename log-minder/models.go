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
	Details map[string]interface{} `json:"details"`
}

type WriteLogPayload struct {
	User    User   `json:"user"`
	Action  Action `json:"action"`
	Message string `json:"message"`
}

type RetreiveLogsFilter struct {
	Limit      int32     `json:"limit" form:"limit"`
	Offset     int32     `json:"offset" form:"offset"`
	Search     string    `json:"search" form:"search"`
	UserID     string    `json:"user_id" form:"user_id"`
	ActionType string    `json:"action_type" form:"action_type"`
	UserRole   string    `json:"user_role" form:"user_role"`
	FromDate   time.Time `json:"from_date"`
	ToDate     time.Time `json:"to_date"`
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
