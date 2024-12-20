package logging

import (
	"context"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

func TestMain(m *testing.M) {
	// Set the AWS_REGION environment variable for tests
	os.Setenv("AWS_REGION", "us-west-2")

	// Run the tests
	code := m.Run()

	// Clean up
	os.Unsetenv("AWS_REGION")

	os.Exit(code)
}

func TestNewCloudWatchLogger(t *testing.T) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		t.Fatalf("unable to load SDK config, %v", err)
	}

	logger := NewCloudWatchLogger(cfg)
	if logger == nil {
		t.Errorf("expected non-nil logger")
	}
}

func TestSendLog(t *testing.T) {
	mockClient := &MockCloudWatchLogsClient{
		PutLogEventsFunc: func(ctx context.Context, params *cloudwatchlogs.PutLogEventsInput, optFns ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.PutLogEventsOutput, error) {
			return &cloudwatchlogs.PutLogEventsOutput{}, nil
		},
	}

	logger := &CloudWatchLogger{
		client:        mockClient,
		logGroupName:  "your-log-group",
		logStreamName: "your-log-stream",
	}

	err := logger.SendLog("test log message")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
