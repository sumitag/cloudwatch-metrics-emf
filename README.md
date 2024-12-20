# CloudWatch Metrics Daemon

A lightweight daemon that collects CPU, memory, GPU, and network metrics, formats them in Embedded Metric Format (EMF), and sends them to AWS CloudWatch Logs.

## Features
- Collects system metrics (CPU, memory, GPU, network).
- Outputs metrics in EMF for seamless integration with CloudWatch.
- Runs as a `systemd` service on Linux or a Windows Service.

## Installation

### Linux
1. Build the application:
   ```bash
   go build -o /usr/local/bin/cloudwatch-metrics-emf
   ```
2. Install the systemd service:
   ```bash
   sudo cp service/systemd.service /etc/systemd/system/cloudwatch-metrics-emf.service
   sudo systemctl enable cloudwatch-metrics-emf
   sudo systemctl start cloudwatch-metrics-emf
   ```

### Windows
1. Build the application:
   ```bash
   go build -o cloudwatch-metrics-emf.exe
   ```
2. Install as a Windows Service:
   Use the Windows Service implementation in `windows_service.go`.

## Configuration
Environment variables for AWS credentials and region are required.

## Logs
Metrics are sent to CloudWatch Logs under the log group `/aws/metrics/system`.
