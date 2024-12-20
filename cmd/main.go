package main

import (
	"cloudwatch-metrics-emf/internal/config"
	"cloudwatch-metrics-emf/internal/logging"
	"cloudwatch-metrics-emf/internal/metrics"
	"log"
	"time"
)

func main() {
	cfg := config.LoadConfig()

	// Load AWS SDK config
	awsCfg, err := config.LoadAWSConfig()
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Initialize CloudWatch logger
	cwLogger := logging.NewCloudWatchLogger(awsCfg)

	// Run the metric collection loop
	for {
		// Collect metrics
		metricData := metrics.CollectMetrics()

		// Convert metricData to EMF JSON
		metricDataStr, err := metricData.ToEMF()
		if err != nil {
			log.Printf("Error converting metrics to EMF: %v\n", err)
			continue
		}

		// Send metrics to CloudWatch Logs
		err = cwLogger.SendLog(metricDataStr)
		if err != nil {
			log.Printf("Error sending log to CloudWatch: %v\n", err)
		}

		// Sleep for the configured interval
		time.Sleep(time.Duration(cfg.Interval) * time.Second)
	}
}
