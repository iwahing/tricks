package utils

import (
	"errors"
	"fmt"
	"time"
)

func GetDaysInSeconds(days int64) int64 {
	return (days * 86400)
}

func GetDaysSinceUnixEpoch(timestamp int64) int64 {
	// Number of days since January 1, 1970, UTC.
	return timestamp / 86400
}

func GetDateTimestamp(datetime string) (int64, error) {
	date, err := time.Parse(time.RFC3339, datetime)
	if err != nil {
		fmt.Println("Error parsing datetime:", err)
		return 0, errors.New("cannot parsing datetime. format must be in RFC3339")
	}

	return date.Unix(), nil
}

func GetDateBefore(dateTimestamp int64, seconds int64) string {
	return UnixTimestampToUTCDateTimeString(dateTimestamp - seconds)
}

func UnixTimestampToUTCDateTimeString(timestamp int64) string {
	// Convert Unix timestamp to time.Time
	t := time.Unix(timestamp, 0)
	// Convert to UTC time
	utcTime := t.UTC()
	// Format the time.Time to the desired layout (RFC3339)
	return utcTime.Format(time.RFC3339)
}
