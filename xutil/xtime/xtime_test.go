package xtime

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"log"
	"testing"
	"time"
)

func TestTimeInUTC(t *testing.T) {
	//loc, err := time.LoadLocation("BST")
	//assert.Equal(t, err, nil)
	//return t.In(loc).Format(format), nil
	//const timeFormat = "2 Jan, 2006 3:04pm (MST)"
	//test , _ := time.Parse( timeFormat, "20 May, 2021 3:08am (BST)" )
	//fmt.Println( test , test.UTC())
	//dur , _ := time.ParseDuration( "1m" )
	//test = test.Add( dur )
	//fmt.Println( test , test.UTC())

	const timeFormat = "2 Jan, 2006 3:04pm (MST)"
	test, err := time.Parse(timeFormat, "25 Oct, 2015 1:59am (BST)")
	fmt.Println(test, test.UTC(), err)
	dur, _ := time.ParseDuration("1m")
	test = test.Add(dur)
	fmt.Println(test, test.UTC())

	fmt.Println("--------------------")

	test, err = time.Parse(timeFormat, "25 Oct, 2015 2:01am (BST)")
	fmt.Println(test, test.UTC(), err)
	test = test.Add(dur)
	fmt.Println(test, test.UTC())

	//t1, err := time.Parse("Jan _2, 2006 | 15:04:05 MST", "May 20, 2021 | 03:08:40 BST")

	//assert.Equal(t, err, nil)
	//log.Println(t1.Unix())
	//assert.Equal(t, t1.Unix(), 0)
}

func TestParseTimeInUTC(t *testing.T) {
	timeLoc, err := ParseTimeInUTC("5/5/2021", "1/2/2006")
	assert.Equal(t, err, nil)
	log.Println(timeLoc.Format("2006-01-02 03:04:05"))
}

func TestIsInOneDay(t *testing.T) {
	is, err := IsInOneDay(1621566534670, 1621567090732)
	assert.Equal(t, err, nil)
	log.Println(is)
}

func TestFirstAndLastTimestampOfMonth(t *testing.T) {
	now := time.Now()
	f, l, err := FirstAndLastTimestampOfMonth(now)
	assert.Equal(t, err, nil)
	log.Println(f.Unix(), l.Unix())
}

func TestParseTimeInLoc(t *testing.T) {
	_, err := ParseTimeInLoc("6 8, 2021 | 08:58:16 BST", "Jan _2, 2006 | 15:04:05 MST", "Europe/London")
	assert.Equal(t, err, nil)
}

func TestGetNextMonthStartEndInUTC(t *testing.T) {
	ts := 1637132342199 / 1000
	tu := time.Unix(int64(ts), 0)
	log.Println(tu)
	startTime, endTime := GetNextMonthStartEndInUTC(tu)
	log.Println(startTime, endTime)
}