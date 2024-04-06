package ds

import "time"

const (
	DateTimeDefaultLayout = "2006-01-02 15:04:05"
)

func DateAfterNow(date string) (bool, error) {
	d, err := time.Parse(DateTimeDefaultLayout, date)
	if err != nil {
		return false, err
	}
	return d.After(time.Now()), nil
}
