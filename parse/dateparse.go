package parse

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Private function to catch Epoch integer in string
func ifEpoch(s string) (time.Time, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	if len(s) >= 10 {

		seconds, err := strconv.ParseInt(s[0:10], 10, 64)
		if err != nil {
			return time.Time{}, err
		}

		nsec, err := strconv.ParseInt(s[10:], 10, 64)
		t := time.Unix(seconds, nsec*1000000)
		return t, err
	}

	t := time.Unix(i, 0)
	return t, err

}

// DateTimeParse -- variety of expected dates
type DateTimeParse string

// getTime --
func (s DateTimeParse) GetTime() (time.Time, error) {

	t, err := ifEpoch(string(s))
	if err == nil {
		return t, err
	}

	layout := []string{
		"January 2, 2006, 3:04 pm",
		"January 2, 2006, 3:04pm",
		"January 2, 2006, 03:04 pm",
		"January 2 2006, 03:04 pm",
		"January 2 2006 03:04 pm",

		"Mon Jan 2  15:04:05 UTC 2006",
		"Mon Jan 02  15:04:05 UTC 2006",

		"January 2, 2006, 3:04 pm",
		"January 2 2006, 3:04 pm",
		"January 2 2006 3:04 pm",
		"Jan 2, 2006, 03:04 pm",
		"Jan 2 2006, 03:04 pm",
		"Jan 2, 2006, 3:04 pm",
		"Jan 2, 06, 3:04 pm",

		"Jan 2  15:04:05",
		"2 Jan 15:04:05",
		"15:04:05 2 Jan",

		"2006-01-02 3:04 pm",
		"2006-01-02 3:04pm",
		"2006-01-02 3:04 PM",
		"2006-01-02 3:04PM",
		"2006-01-02 15:04",

		"2006/01/02 3:04 pm",
		"2006/01/02 3:04pm",
		"2006/01/02 3:04 PM",
		"2006/01/02 3:04PM",
		"2006/01/02 15:04",

		"01/02/2006 3:04 pm",
		"01/02/2006 3:04pm",
		"01/02/2006 3:04 PM",
		"01/02/2006 3:04PM",
		"01/02/2006 15:04",

		"01.02.2006 3:04 pm",
		"01.02.2006 3:04pm",
		"01.02.2006 3:04 PM",
		"01.02.2006 3:04PM",
		"01.02.2006 15:04",

		"01/02/2006",
		"1/2/2006",

		"01_02_2006",
		"1_2_2006",

		"01-02-2006",
		"1-2-2006",

		"01.02.2006",
		"1.2.2006",

		"1.2.2006  15:04",
		"1/2/2006  15:04",
		"1_2_2006  15:04",
		"1 2 2006  15:04",

		"15:04 1.2.2006",
		"15:04 1/2/2006",
		"15:04 1_2_2006",
		"15:04 1 2 2006",

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

/*  NewYork() - Input localtime New_York and convert to UTC
          Add year, if missing

    Input: "Sep  8  13:24:18 "
    Expected output: "2018-09-08 13:24:18 -0400 EDT"


*/
func (s DateTimeParse) NewYork() (time.Time, error) {

	tt, err := DateTimeParse(s).GetTime()
	if err != nil {
		return tt, err
	}

	if tt.Year() == 0 {
		tt = tt.AddDate(time.Now().Year(), 0, 0)
	}

	loc, err := time.LoadLocation("America/New_York")

	_, offset := tt.In(loc).Zone()
	tt = tt.Add(time.Duration(-offset) * time.Second)

	return tt.In(loc), err
}

// GetTimeLoc --
func (s DateTimeParse) GetTimeLoc() (time.Time, error) {

	tt, err := DateTimeParse(s).GetTime()
	if err != nil {
		return tt, err
	}
	loc, err := time.LoadLocation("America/New_York")

	return tt.In(loc), err

}

// GetTimeLocSquish -- Force min to be int in 10 min interval
func (s DateTimeParse) GetTimeLocSquish() (string, error) {

	tt, err := DateTimeParse(s).GetTime()
	if err != nil {
		return "", err
	}
	squishMin := int(tt.Minute()/10) * 10
	ret := fmt.Sprintf("%02d:%02d", tt.Hour(), squishMin)
	return ret, err

}

// GetTimeLocHRminS --
func (s DateTimeParse) GetTimeLocHRminS() (string, error) {

	tt, err := DateTimeParse(s).GetTime()
	if err != nil {
		return "", err
	}
	ret := fmt.Sprintf("%02d:%02d", tt.Hour(), tt.Minute())
	return ret, err

}

func (s DateTimeParse) DaysFrom(day2 string) (int, error) {

	tt, err := DateTimeParse(s).GetTime()
	if err != nil {
		return 0, err
	}

	t2, err := DateTimeParse(day2).GetTime()
	if err != nil {
		return 0, err
	}

	days := int(t2.Sub(tt).Hours() / 24)

	return days, err

}

// DaysBetween always positive
func (s DateTimeParse) DaysBetween(day2 string) (int, error) {

	tt, err := DateTimeParse(s).GetTime()
	if err != nil {
		return 0, err
	}

	t2, err := DateTimeParse(day2).GetTime()
	if err != nil {
		return 0, err
	}

	days := int(t2.Sub(tt).Hours() / 24)
	if days < 0 {
		days = -days
	}

	return days, err

}
