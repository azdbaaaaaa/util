package xtime

import (
	"time"
)

const (
	DateFormat              = "2006-01-02"
	DateTimeFormat          = "2006-01-02 15:04:05"
	DateTimeFormat_YYYYMMDD = "20060102"
	DateTimeFormat_YYYYMM   = "200601"
	DateTimeFormat_YYYY_MM  = "2006-01"
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

func YesterdayInUTC() time.Time {
	currentTime := time.Now().UTC()
	return currentTime.AddDate(0, 0, -1)
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

func GetYesterdayStartTimeInUTC() int64 {
	t := time.Now().UTC()
	yesterdayDay := t.AddDate(0, 0, -1)
	return time.Date(yesterdayDay.Year(), yesterdayDay.Month(), yesterdayDay.Day(), 0, 0, 0, 0, yesterdayDay.Location()).UnixNano() / 1e6
}

func GetYesterdayEndTimeInUTC() int64 {
	t := time.Now().UTC()
	yesterdayDay := t.AddDate(0, 0, -1)
	return time.Date(yesterdayDay.Year(), yesterdayDay.Month(), yesterdayDay.Day(), 23, 59, 59, 0, yesterdayDay.Location()).UnixNano() / 1e6
}

func GetNextMonthStartEndInUTC(t time.Time) (time.Time, time.Time) {
	t = t.UTC()
	monthStartDay := t.AddDate(0, 1, -t.Day()+1)
	monthStartTime := time.Date(monthStartDay.Year(), monthStartDay.Month(), monthStartDay.Day(), 0, 0, 0, 0, t.Location())
	monthEndDay := monthStartTime.AddDate(0, 1, -1)
	monthEndTime := time.Date(monthEndDay.Year(), monthEndDay.Month(), monthEndDay.Day(), 23, 59, 59, 0, t.Location())
	return monthStartTime, monthEndTime
}

func ParseStringToInt(date string) int64 {
	t, _ := time.ParseInLocation(DateTimeFormat_YYYY_MM, date, time.UTC)
	return t.UnixNano() / 1e6
}

func CalAge(birthdate, today time.Time) int {
	today = today.In(birthdate.Location())
	ty, tm, td := today.Date()
	today = time.Date(ty, tm, td, 0, 0, 0, 0, time.UTC)
	by, bm, bd := birthdate.Date()
	birthdate = time.Date(by, bm, bd, 0, 0, 0, 0, time.UTC)
	if today.Before(birthdate) {
		return 0
	}
	age := ty - by
	anniversary := birthdate.AddDate(age, 0, 0)
	if anniversary.After(today) {
		age--
	}
	return age
}

func IsAdult(layout, birth string) (bool, error) {
	birthday, err := time.Parse(layout, birth)
	if err != nil {
		return false, err
	}
	age := CalAge(birthday, time.Now())
	if age >= 18 {
		return true, nil
	}
	return false, nil
}
