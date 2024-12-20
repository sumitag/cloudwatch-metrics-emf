package logging

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
)

type CloudWatchLogsClient interface {
	PutLogEvents(ctx context.Context, params *cloudwatchlogs.PutLogEventsInput, optFns ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.PutLogEventsOutput, error)
}

type CloudWatchLogger struct {
	client        CloudWatchLogsClient
	logGroupName  string
	logStreamName string
}

func NewCloudWatchLogger(cfg aws.Config) *CloudWatchLogger {
	client := cloudwatchlogs.NewFromConfig(cfg)

	return &CloudWatchLogger{
		client:        client,
		logGroupName:  "your-log-group",
		logStreamName: "your-log-stream",
	}
}

func (l *CloudWatchLogger) SendLog(message string) error {
	input := &cloudwatchlogs.PutLogEventsInput{
		LogGroupName:  &l.logGroupName,
		LogStreamName: &l.logStreamName,
		LogEvents: []types.InputLogEvent{
			{
				Timestamp: aws.Int64(time.Now().UnixNano() / int64(time.Millisecond)),
				Message:   aws.String(message),
			},
		},
	}

	_, err := l.client.PutLogEvents(context.TODO(), input)
	return err
}
