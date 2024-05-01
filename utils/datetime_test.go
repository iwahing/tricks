package utils

import (
	"testing"
	"time"
)

func TestGetDaysInSeconds(t *testing.T) {
	daysInSeconds := GetDaysInSeconds(9)
	if daysInSeconds != int64(777_600) {
		t.Errorf("GetDaysInSeconds(9) = %d; want 777600", daysInSeconds)
	}
}

func TestGetDaysSinceUnixEpoch(t *testing.T) {
	timestamp := int64(1714534078)
	daysInSeconds := GetDaysSinceUnixEpoch(timestamp)
	if daysInSeconds != int64(19844) {
		t.Errorf("GetDaysSinceUnixEpoch(%d) = %d; want 19844", daysInSeconds, timestamp)
	}
}

func TestGetDateTimestamp(t *testing.T) {
	datetime_string := "2024-04-30T12:00:00Z"
	timestamp, _ := GetDateTimestamp(datetime_string)

	date, _ := time.Parse(time.RFC3339, datetime_string)
	correctValue := date.Unix()
	if timestamp != correctValue {
		t.Errorf("GetDateTimestamp(%s) = %d; want %d", datetime_string, timestamp, correctValue)
	}
}

func TestGetDateTimestampWhenError(t *testing.T) {
	datetime_string := "2024-04-30 12:00:00Z"
	_, err := GetDateTimestamp(datetime_string)

	if err == nil {
		t.Errorf("GetDateTimestamp(%s) Error Parsing data should fail. Wrong format", datetime_string)
	}
}

func TestGetDateBefore(t *testing.T) {
	fiveDaysInSeconds := int64(5 * 24 * 60 * 60)
	initialDateTimestamp := int64(1714478400)
	daysInSeconds := GetDateBefore(initialDateTimestamp, fiveDaysInSeconds)

	timeT := time.Unix(initialDateTimestamp, 0)
	fiveDaysAgo := timeT.AddDate(0, 0, -5)
	utcTime := fiveDaysAgo.UTC()
	timestampFiveDaysAgo := utcTime.Format(time.RFC3339)

	if daysInSeconds != timestampFiveDaysAgo {
		t.Errorf("GetDateBefore(%d) = %s; want %s", initialDateTimestamp, daysInSeconds, timestampFiveDaysAgo)
	}
}

func TestUnixTimestampToUTCDateTimeString(t *testing.T) {
	initialDateTimestamp := int64(1714017600)
	datetime_string := UnixTimestampToUTCDateTimeString(initialDateTimestamp)

	actualValue := "2024-04-25T04:00:00Z"
	if datetime_string != actualValue {
		t.Errorf("UnixTimestampToUTCDateTimeString(%d) = %s; want %s", initialDateTimestamp, datetime_string, actualValue)
	}
}
