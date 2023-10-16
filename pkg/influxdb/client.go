package influxdb

import (
	"context"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

const (
	influxToken  = "INFLUX_DB_TOKEN"
	influxOrg    = "iscsi"
	influxBucket = "iscsi_performance"
	influxURL    = "http://localhost:8086"
)

type Client struct {
	client   influxdb2.Client
	writeAPI api.WriteAPIBlocking
}

// NewClient initializes and returns a new InfluxDB client.
func NewClient() *Client {
	client := influxdb2.NewClient(influxURL, influxToken)
	writeAPI := client.WriteAPIBlocking(influxOrg, influxBucket)

	return &Client{
		client:   client,
		writeAPI: writeAPI,
	}
}

// Close closes the InfluxDB client connection.
func (c *Client) Close() {
	c.client.Close()
}

// WriteData writes performance data to InfluxDB.
func (c *Client) WriteData(operation string, fileSize int, duration time.Duration, speed float64) {
	p := influxdb2.NewPoint(
		"iscsi_performance",
		map[string]string{"operation": operation},
		map[string]interface{}{
			"file_size": fileSize,
			"duration":  duration.Seconds(),
			"speed":     speed,
		},
		time.Now(),
	)
	err := c.writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		fmt.Printf("Error writing to InfluxDB: %v\n", err)
	}
	c.client.Close()
}
