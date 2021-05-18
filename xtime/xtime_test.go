package xtime

import (
	"github.com/go-playground/assert/v2"
	"log"
	"testing"
)

func TestTimeInUTC(t *testing.T) {
	
}

func TestParseTimeInUTC(t *testing.T) {
	timeLoc, err := ParseTimeInUTC("5/5/2021", "1/2/2006")
	assert.Equal(t, err, nil)
	log.Println(timeLoc.Format("2006-01-02 03:04:05"))
}