package xtime

import "time"

const (
	DateFormat     = "2006-01-02"
	DateTimeFormat = "2006-01-02 15:04:05"
)

func TimeInUTC(t time.Time, format string) (string, error) {
	// https:/golang.org/pkg/xtime/#LoadLocation loads location on
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		return "", err
	}
	return t.In(loc).Format(format), nil
}

func ParseTimeInUTC(t, format string) (time.Time, error) {
	tt, err := time.Parse(format, t)
	if err != nil {
		return time.Time{}, err
	}
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		return time.Time{}, err
	}
	return tt.In(loc), nil
}
