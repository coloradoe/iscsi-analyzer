package main

import (
	"fmt"
	"os"
	"time"
)

const (
	testFilePath = "/mnt/iscsi_target/testfile" // Adjust this path to your iSCSI mount point
	testFileSize = 100 * 1024 * 1024            // 100 MB
)

func writeTest() time.Duration {
	start := time.Now()

	data := make([]byte, testFileSize)
	err := os.WriteFile(testFilePath, data, 0644)
	if err != nil {
		fmt.Println("Error writing test file:", err)
		return 0
	}

	return time.Since(start)
}

func readTest() time.Duration {
	start := time.Now()

	_, err := os.ReadFile(testFilePath)
	if err != nil {
		fmt.Println("Error reading test file:", err)
		return 0
	}

	return time.Since(start)
}

func main() {
	// Write Test
	writeDuration := writeTest()
	if writeDuration != 0 {
		fmt.Printf("Write speed: %.2f MB/s\n", float64(testFileSize)/writeDuration.Seconds()/(1024*1024))
	}

	// Read Test
	readDuration := readTest()
	if readDuration != 0 {
		fmt.Printf("Read speed: %.2f MB/s\n", float64(testFileSize)/readDuration.Seconds()/(1024*1024))
	}
}
