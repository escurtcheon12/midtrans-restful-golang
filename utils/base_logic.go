package utils

import (
	"log"
	"time"
)

func ParseTimestamp(timestampStr string) time.Time {
	if timestampStr == "" {
		log.Println("Warning: Timestamp is an empty string")
		return time.Time{}
	}

	timestamp, err := time.Parse("2006-01-02 15:04:05", timestampStr)
	if err != nil {
		log.Printf("Error parsing timestamp: %v", err)
		return time.Time{}
	}

	return timestamp.UTC()
}
