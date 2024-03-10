package main

import (
	"context"
)

type LogMinderClient interface {
	WriteLog(ctx context.Context, payload *WriteLogPayload) error
	RetreiveLogs(ctx context.Context, payload *RetreiveLogsFilter) (*RetreiveLogsResponse, error)
}
