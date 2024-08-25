package datetimes

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_FormatTimeAsAMonthDayString(t *testing.T) {
	datetime := time.UnixMilli(11000222222)
	expected := "5/8"
	actual := FormatDatetimeToMMDD(&datetime)

	assert.Equal(
		t, expected, actual,
		fmt.Sprintf("Actual [%s] did not match expected [%s]", actual, expected),
	)
}

func Test_GetCurrentDateInMMDD(t *testing.T) {
	now := time.Now()
	expected := FormatDatetimeToMMDD(&now)
	actual := GetCurrentDateInMMDD()

	assert.Equal(
		t, expected, actual,
		fmt.Sprintf("Actual [%s] did not match expected [%s]", actual, expected),
	)
}

func Test_GetCurrentDateMinus1InMMDD(t *testing.T) {
	now := time.Now().AddDate(0, 0, -1)
	expected := FormatDatetimeToMMDD(&now)
	actual := GetCurrentDateMinus1InMMDD()

	assert.Equal(
		t, expected, actual,
		fmt.Sprintf("Actual [%s] did not match expected [%s]", actual, expected),
	)
}

func Test_GetCurrentDateMinus7InMMDD(t *testing.T) {
	now := time.Now().AddDate(0, 0, -7)
	expected := FormatDatetimeToMMDD(&now)
	actual := GetCurrentDateMinus7InMMDD()

	assert.Equal(
		t, expected, actual,
		fmt.Sprintf("Actual [%s] did not match expected [%s]", actual, expected),
	)
}

func Test_FormatDatetimeToTimezone(t *testing.T) {
	datetime := time.UnixMilli(11000222222)

	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Fatalf("Error loading location")
	}

	expected := "08-May-1970 3:37"
	actual := FormatDatetimeToET(&datetime, location)

	assert.Equal(
		t, expected, actual,
		fmt.Sprintf("Actual [%s] did not match expected [%s]", actual, expected),
	)
}
