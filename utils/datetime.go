package utils

import (
	"fmt"
	"time"
)

// CustomError is a custom error type
type CustomError struct {
	message string
}

// Error returns the error message
func (e *CustomError) Error() string {
	return e.message
}

func GetSeconds(days int64) int64 {
	return (days * 3600)
}

func GetDaysSinceUnixEpoch(timestamp int64) int64 {
	// Number of days since January 1, 1970, UTC.
	return timestamp / 86400
}

func GetDateTimestamp(datetime string) (int64, error) {
	date, err := time.Parse(time.RFC3339, datetime)
	if err != nil {
		fmt.Println("Error parsing datetime:", err)
		return 0, &CustomError{"Error parsing datetime:"}
	}

	return date.Unix(), nil
}

func GetDateBefore(dateTimestamp int64, days_in_second int64) int64 {
	return dateTimestamp - days_in_second
}

func UnixTimestampToUTCDateTimeString(timestamp int64) string {
	// Convert Unix timestamp to time.Time
	t := time.Unix(timestamp, 0)
	// Convert to UTC time
	utcTime := t.UTC()
	// Format the time.Time to the desired layout (RFC3339)
	return utcTime.Format(time.RFC3339)
}
