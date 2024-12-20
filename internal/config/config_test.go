package config

import (
	"os"
	"testing"
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

func TestLoadConfig(t *testing.T) {
	cfg := LoadConfig()
	if cfg.Interval != 60 {
		t.Errorf("expected Interval to be 60, got %d", cfg.Interval)
	}
}

func TestLoadAWSConfig(t *testing.T) {
	cfg, err := LoadAWSConfig()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if cfg.Region != "us-west-2" {
		t.Errorf("expected region to be us-west-2, got %v", cfg.Region)
	}
}
