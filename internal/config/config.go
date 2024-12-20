package config

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

type Config struct {
	Interval int
}

func LoadConfig() *Config {
	// Load default AWS configuration
	_, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return &Config{
		Interval: 60, // Default collection interval (seconds)
	}
}

// LoadAWSConfig loads the AWS SDK configuration
func LoadAWSConfig() (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("unable to load SDK config, %v", err)
		return aws.Config{}, err
	}
	return cfg, nil
}
