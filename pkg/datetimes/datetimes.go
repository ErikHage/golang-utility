package datetimes

import (
	"fmt"
	"time"
)

const (
	DatetimeFormat = "02-Jan-2006 3:04"
)

func FormatDatetimeToMMDD(datetime *time.Time) string {
	if datetime == nil {
		return ""
	}

	return fmt.Sprintf("%d/%d", int(datetime.Month()), datetime.Day())
}

func GetCurrentDateInMMDD() string {
	datetime := time.Now()

	return FormatDatetimeToMMDD(&datetime)
}

func GetCurrentDateMinus1InMMDD() string {
	datetime := time.Now()
	datetime = datetime.AddDate(0, 0, -1)

	return FormatDatetimeToMMDD(&datetime)
}

func GetCurrentDateMinus7InMMDD() string {
	datetime := time.Now()
	datetime = datetime.AddDate(0, 0, -7)

	return FormatDatetimeToMMDD(&datetime)
}

func FormatDatetimeToET(datetime *time.Time, location *time.Location) string {
	if datetime == nil || location == nil {
		return ""
	}

	return datetime.In(location).Format(DatetimeFormat)
}
