package dateparse

import (
	"fmt"
	"strings"
	"time"
)

// DateTimeParse -- variety of expected dates
type DateTimeParse string

// getTime --
func (s DateTimeParse) getTime() (time.Time, error) {
	layout := []string{
		"January 2, 2006, 3:04 pm",
		"January 2, 2006, 3:04pm",
		"January 2, 2006, 03:04 pm",
		"January 2 2006, 03:04 pm",
		"January 2 2006 03:04 pm",

		"January 2, 2006, 3:04 pm",
		"January 2 2006, 3:04 pm",
		"January 2 2006 3:04 pm",
		"Jan 2, 2006, 03:04 pm",
		"Jan 2 2006, 03:04 pm",
		"Jan 2, 2006, 3:04 pm",
		"Jan 2, 06, 3:04 pm",
		"2006-01-02 3:04 pm",
		"2006-01-02 3:04pm",
		"2006-01-02 3:04 PM",
		"2006-01-02 3:04PM",
		"2006-01-02 15:04",

		"2006-01-02 15:04:05 +0000 UTC",
		"2006-01-02T15:04:05+07:00",
		"2006-01-02 15:04:05 +07:00",
	}

	st := strings.Join(strings.Fields(string(s)), " ")
	//fmt.Printf("-->%s\n", st)

	for _, l := range layout[:len(layout)-1] {
		t, err := time.Parse(l, st)
		if err == nil {
			return t, err
		}

	}

	return time.Parse(layout[len(layout)-1], st)

}

// getTimeLoc --
func (s DateTimeParse) getTimeLoc() (time.Time, error) {

	tt, err := DateTimeParse(s).getTime()
	if err != nil {
		return tt, err
	}
	loc, err := time.LoadLocation("America/New_York")

	return tt.In(loc), err

}

func (s DateTimeParse) getTimeLocSquish() (string, error) {

	tt, err := DateTimeParse(s).getTime()
	if err != nil {
		return "", err
	}
	squishMin := int(tt.Minute()/10) * 10
	ret := fmt.Sprintf("%02d:%02d", tt.Hour(), squishMin)
	return ret, err

}

func (s DateTimeParse) getTimeLocHRminS() (string, error) {

	tt, err := DateTimeParse(s).getTime()
	if err != nil {
		return "", err
	}
	ret := fmt.Sprintf("%02d:%02d", tt.Hour(), tt.Minute())
	return ret, err

}
