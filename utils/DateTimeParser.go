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

// GetDateRange converts a predefined date range string into a start and end time.
func GetDateRange(dateRange string) (start time.Time, end time.Time) {
	now := time.Now()

	switch dateRange {
	case "today":
		start = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		end = start.Add(24 * time.Hour).Add(-time.Second)
	case "tomorrow":
		start = time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
		end = start.Add(24 * time.Hour).Add(-time.Second)
	case "thisWeek":
		weekday := int(now.Weekday())
		if weekday == 0 { // Adjust if today is Sunday
			weekday = 7
		}
		start = now.AddDate(0, 0, -weekday+1) // Start of the week (Monday)
		start = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())
		end = start.AddDate(0, 0, 6) // End of the week (Sunday)
		end = time.Date(end.Year(), end.Month(), end.Day(), 23, 59, 59, 999, end.Location())
	case "thisMonth":
		start = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		end = start.AddDate(0, 1, -1) // Last day of the month
		end = time.Date(end.Year(), end.Month(), end.Day(), 23, 59, 59, 999, end.Location())
	case "nextWeek":
		nextMonday := now.AddDate(0, 0, 8-int(now.Weekday())) // Next week's Monday
		start = time.Date(nextMonday.Year(), nextMonday.Month(), nextMonday.Day(), 0, 0, 0, 0, nextMonday.Location())
		end = start.AddDate(0, 0, 6) // Next week's Sunday
		end = time.Date(end.Year(), end.Month(), end.Day(), 23, 59, 59, 999, end.Location())
	case "nextMonth":
		start = time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, now.Location())
		end = start.AddDate(0, 1, -1) // Last day of next month
		end = time.Date(end.Year(), end.Month(), end.Day(), 23, 59, 59, 999, end.Location())
	default:
		// If no range is provided, return zero values (no filter applied)
		return time.Time{}, time.Time{}
	}

	return start, end
}
