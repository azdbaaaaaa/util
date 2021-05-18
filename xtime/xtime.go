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

func IsInOneDay(ts1, ts2 int64) (equal bool, err error) {
	t1, err := TimeInUTC(time.Unix(ts1, 0), DateFormat)
	if err != nil {
		return
	}
	t2, err := TimeInUTC(time.Unix(ts2, 0), DateFormat)
	if err != nil {
		return
	}
	return t1 == t2, nil
}
