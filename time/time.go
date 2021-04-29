package time

import "time"

func TimeInUTC(t time.Time, format string) (string, error) {
	// https:/golang.org/pkg/time/#LoadLocation loads location on
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		return "", err
	}
	return t.In(loc).Format(format), nil
}
