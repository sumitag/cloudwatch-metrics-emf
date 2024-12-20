package logging

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

type MockCloudWatchLogsClient struct {
	cloudwatchlogs.Client
	PutLogEventsFunc func(ctx context.Context, params *cloudwatchlogs.PutLogEventsInput, optFns ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.PutLogEventsOutput, error)
}

func (m *MockCloudWatchLogsClient) PutLogEvents(ctx context.Context, params *cloudwatchlogs.PutLogEventsInput, optFns ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.PutLogEventsOutput, error) {
	return m.PutLogEventsFunc(ctx, params, optFns...)
}
