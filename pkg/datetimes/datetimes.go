package datetimes

import (
	"fmt"
	"time"
)

const (
	DatetimeFormat = "02-Jan-2006 3:04"
	NewYork        = "America/New_York"
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

func FormatDatetimeToET(datetime *time.Time) string {
	if datetime == nil {
		return ""
	}

	location, err := time.LoadLocation(NewYork)
	if err != nil {
		return datetime.Format(DatetimeFormat)
	}

	return fmt.Sprintf("%s ET", datetime.In(location).Format(DatetimeFormat))
}
