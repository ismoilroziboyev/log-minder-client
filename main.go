package main

import (
	"context"
	"fmt"
)

type LogMinderClient interface {
	WriteLog(ctx context.Context, payload *WriteLogPayload) error
	RetreiveLogs(ctx context.Context, payload *RetreiveLogsFilter) (*RetreiveLogsResponse, error)
}

func main() {
	fmt.Println("hello world!")
}
