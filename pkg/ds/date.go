package ds

import (
	"log"
	"time"
)

const (
	DateTimeDefaultLayout = "2006-01-02 15:04:05"
)

var (
	DefaultUpstreamTimezone       *time.Location
	DefaultUpstreamTimezoneString = "Asia/Taipei"
)

func DateParse(date string) (time.Time, error) {
	return time.ParseInLocation(DateTimeDefaultLayout, date, DefaultUpstreamTimezone)
}

func DateAfterNow(date string) (bool, error) {
	d, err := time.ParseInLocation(DateTimeDefaultLayout, date, DefaultUpstreamTimezone)
	if err != nil {
		return false, err
	}
	return d.After(time.Now()), nil
}

func init() {
	tz, err := time.LoadLocation(DefaultUpstreamTimezoneString)
	if err != nil {
		log.Fatalln("failed to parse tz", err)
	}
	DefaultUpstreamTimezone = tz
}
