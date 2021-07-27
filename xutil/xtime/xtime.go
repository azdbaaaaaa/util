package xtime

import (
	"time"
)

const (
	DateFormat              = "2006-01-02"
	DateTimeFormat          = "2006-01-02 15:04:05"
	DateTimeFormat_YYYYMMDD = "20060102"
	DateTimeFormat_YYYYMM   = "200601"
)

// FirstAndLastTimestampOfMonth 获取t时间所在月的开始和结束时间
func FirstAndLastTimestampOfMonth(t time.Time) (first, last time.Time, err error) {
	currentYear, currentMonth, _ := t.Date()
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		return
	}
	//time.Time格式
	first = time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, loc)
	last = first.AddDate(0, 1, 0).Add(-1 * time.Second)
	return
}

func TimeInUTC(t time.Time, format string) (string, error) {
	// https:/golang.org/pkg/xtime/#LoadLocation loads location on
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		return "", err
	}
	return t.In(loc).Format(format), nil
}

func ParseTimeInLoc(value, format, locName string) (t time.Time, err error) {
	if locName == "" {
		locName = "Europe/London"
	}
	loc, err := time.LoadLocation(locName)
	if err != nil {
		return time.Time{}, err
	}
	t, err = time.ParseInLocation(format, value, loc)
	if err != nil {
		return time.Time{}, err
	}
	return
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
	t1, err := TimeInUTC(time.Unix(ts1/1000, 0), DateFormat)
	if err != nil {
		return
	}
	t2, err := TimeInUTC(time.Unix(ts2/1000, 0), DateFormat)
	if err != nil {
		return
	}
	return t1 == t2, nil
}

func TodayInUTC() time.Time {
	return time.Now().UTC()
}

func GetMonthStartEndInUTC(t time.Time) (time.Time, time.Time) {
	t = t.UTC()
	monthStartDay := t.AddDate(0, 0, -t.Day()+1)
	monthStartTime := time.Date(monthStartDay.Year(), monthStartDay.Month(), monthStartDay.Day(), 0, 0, 0, 0, t.Location())
	monthEndDay := monthStartTime.AddDate(0, 1, -1)
	monthEndTime := time.Date(monthEndDay.Year(), monthEndDay.Month(), monthEndDay.Day(), 23, 59, 59, 0, t.Location())
	return monthStartTime, monthEndTime
}

func GetLastMonthStartEndInUTC(t time.Time) (time.Time, time.Time) {
	t = t.UTC()
	monthStartDay := t.AddDate(0, -1, -t.Day()+1)
	monthStartTime := time.Date(monthStartDay.Year(), monthStartDay.Month(), monthStartDay.Day(), 0, 0, 0, 0, t.Location())
	monthEndDay := monthStartTime.AddDate(0, 1, -1)
	monthEndTime := time.Date(monthEndDay.Year(), monthEndDay.Month(), monthEndDay.Day(), 23, 59, 59, 0, t.Location())
	return monthStartTime, monthEndTime
}

func GetTodayStartTimeInUTC() int64 {
	currentTime := time.Now().UTC()
	return time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location()).UnixNano() / 1e6
}

func GetTodayEndTimeInUTC() int64 {
	currentTime := time.Now().UTC()
	return time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location()).UnixNano() / 1e6
}