package utils

import "time"

func DateParser(dateStr string) time.Time {
	parsedDate, _ := time.Parse("2006-01-02", dateStr)
	return parsedDate
}
func TimeParser(timeStr string) time.Time {
	parsedTime, _ := time.Parse("15:04:05", timeStr)
	return parsedTime
}