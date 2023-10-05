package iscsi

import (
	"fmt"
	"os"
	"time"
)

const (
	testFilePath = "/mnt/iscsi_target/testfile"
)

// ReadTest measures the time taken to read the test file from the iSCSI target.
func ReadTest(fileSize int) time.Duration {
	start := time.Now()

	_, err := os.ReadFile(testFilePath)
	if err != nil {
		fmt.Println("Error reading test file:", err)
		return 0
	}

	return time.Since(start)
}
