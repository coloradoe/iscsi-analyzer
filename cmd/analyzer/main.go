package main

import (
	"fmt"

	"github.com/coloradoe/iscsi-analyzer/pkg/influxdb"
	"github.com/coloradoe/iscsi-analyzer/pkg/iscsi"
)

const (
	testFileSize = 100 * 1024 * 1024 // 100 MB
)

func main() {
	// Initialize InfluxDB client
	influxClient := influxdb.NewClient()
	defer influxClient.Close()

	// Write Test
	writeDuration := iscsi.WriteTest(testFileSize)
	if writeDuration != 0 {
		speed := float64(testFileSize) / writeDuration.Seconds() / (1024 * 1024)
		fmt.Printf("Write speed: %.2f MB/s\n", speed)
		influxClient.WriteData("write", testFileSize, writeDuration, speed)
	}

	// Read Test
	readDuration := iscsi.ReadTest(testFileSize)
	if readDuration != 0 {
		speed := float64(testFileSize) / readDuration.Seconds() / (1024 * 1024)
		fmt.Printf("Read speed: %.2f MB/s\n", speed)
		influxClient.WriteData("read", testFileSize, readDuration, speed)
	}
}
