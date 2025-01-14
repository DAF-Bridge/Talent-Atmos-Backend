package utils

import (
	"fmt"
	"time"
)

func DateParser(dateStr string) time.Time {
	parsedDate, _ := time.Parse("2006-01-02", dateStr)
	return parsedDate
}

func TimeParser(timeStr string) time.Time {
	parsedTime, _ := time.Parse("15:04:05", timeStr)
	return parsedTime
}

func ISO8601Parser(isoStr string) time.Time {
	parsedTime, _ := time.Parse(time.RFC3339, isoStr)
	return parsedTime
}

type DateOnly struct {
	time.Time
}

func (d DateOnly) MarshalJSON() ([]byte, error) {
	// Format the date as "YYYY-MM-DD"
	formatted := fmt.Sprintf(`"%s"`, d.Format("2006-01-02"))
	return []byte(formatted), nil
}

func (d *DateOnly) UnmarshalJSON(b []byte) error {
	// Parse date in "YYYY-MM-DD" format
	parsedDate, err := time.Parse(`"2006-01-02"`, string(b))
	if err != nil {
		return err
	}
	d.Time = parsedDate
	return nil
}
