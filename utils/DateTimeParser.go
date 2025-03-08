package utils

import (
	"database/sql/driver"
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

type TimeOnly struct {
	time.Time
}

func (t TimeOnly) MarshalJSON() ([]byte, error) {
	// Format the time as "HH:MM"
	formatted := fmt.Sprintf(`"%s"`, t.Format("15:04:05"))
	return []byte(formatted), nil
}

func (t *TimeOnly) UnmarshalJSON(b []byte) error {
	// Parse time in "HH:MM" format
	parsedTime, err := time.Parse(`"15:04:05"`, string(b))
	if err != nil {
		return err
	}
	t.Time = parsedTime
	return nil
}

// Value Convert to a database value (for GORM)
func (t TimeOnly) Value() (driver.Value, error) {
	return t.Format("15:04:05"), nil
}

// Convert from database value to struct
//
//	func (t *TimeOnly) Scan(value interface{}) error {
//		if v, ok := value.(time.Time); ok {
//			t.Time = v
//			return nil
//		}
//		return fmt.Errorf("cannot convert %v to TimeOnly", value)
//	}
//
// Ensure TimeOnly implements sql.Scanner
func (t *TimeOnly) Scan(value interface{}) error {
	if value == nil {
		*t = TimeOnly{}
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		*t = TimeOnly{Time: v}
		return nil
	case string:
		parsedTime, err := time.Parse("15:04:05", v)
		if err != nil {
			return fmt.Errorf("cannot parse time: %v", err)
		}
		*t = TimeOnly{Time: parsedTime}
		return nil
	default:
		return fmt.Errorf("cannot convert %T to TimeOnly", value)
	}
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

func (d DateOnly) Value() (driver.Value, error) {
	return d.Format("2006-01-02"), nil
}

//	func (d *DateOnly) Scan(value interface{}) error {
//		if v, ok := value.(time.Time); ok {
//			d.Time = v
//			return nil
//		}
//		return fmt.Errorf("cannot convert %v to DateOnly", value)
//	}
func (d *DateOnly) Scan(value interface{}) error {
	if value == nil {
		*d = DateOnly{}
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		*d = DateOnly{Time: v}
		return nil
	case string:
		parsedDate, err := time.Parse("2006-01-02", v)
		if err != nil {
			return fmt.Errorf("cannot parse date: %v", err)
		}
		*d = DateOnly{Time: parsedDate}
		return nil
	default:
		return fmt.Errorf("cannot convert %T to DateOnly", value)
	}
}

// GetDateRange Searching Service Utils
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
