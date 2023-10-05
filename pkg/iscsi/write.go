package iscsi

import (
	"fmt"
	"os"
	"time"
)

// WriteTest measures the time taken to write a test file to the iSCSI target.
func WriteTest(fileSize int) time.Duration {
	start := time.Now()

	data := make([]byte, fileSize)
	err := os.WriteFile(testFilePath, data, 0644)
	if err != nil {
		fmt.Println("Error writing test file:", err)
		return 0
	}

	return time.Since(start)
}
