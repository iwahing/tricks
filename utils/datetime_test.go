package utils

import (
	"testing"
)

func TestGetDaysInSeconds(t *testing.T) {
	daysInSeconds := GetSeconds(9)
	if daysInSeconds != 32400 {
		t.Errorf("GetSeconds(9) = %d; want 32400", daysInSeconds)
	}
}

// TODO GetDaysSinceUnixEpoch
// TODO GetDateTimestamp
// TODO GetDateBefore
// TODO UnixTimestampToUTCDateTimeString
