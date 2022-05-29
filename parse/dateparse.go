package parse

import (
	"errors"
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

	if len(s) > 10 {

		seconds, _ := strconv.ParseInt(s[0:10], 10, 64)
		nsec, err := strconv.ParseInt(s[10:], 10, 64)

		t := time.Unix(seconds, nsec*1000000)
		return t, err
	}

	t := time.Unix(i, 0)
	return t, err

}

// DateTimeParse -- variety of expected dates
type DateTimeParse string

var lastused = "2006-01-02T15:04:05+07:00"
var layout = []string{

	"January 2 2006 3:04 pm",
	"January 2 2006 3:04pm",
	"January 2 2006 03:04 pm",
	"January 2 2006 03:04 pm",
	"January 2 2006 03:04 pm",

	"January 2 2006 3:04:05 pm",
	"January 2 2006 3:04:05pm",
	"January 2 2006 03:04:05 pm",
	"January 2 2006 03:04:05 pm",
	"January 2 2006 03:04:05 pm",

	"Mon Jan _2  15:04 UTC 2006",
	"Mon Jan _2  15:04:05 UTC 2006",
	"Mon Jan _2  15:04:05 2006",
	"Mon 2006 Jan _2 15:04:05",

	"Mon _2 Jan 15:04 UTC 2006",
	"Mon _2 Jan 15:04:05 UTC 2006",
	"Mon _2 Jan 15:04:05 2006",
	"Mon 2006 Jan _2 15:04:05",

	"Mon January _2  15:04 UTC 2006",
	"Mon January _2  15:04:05 UTC 2006",
	"Mon January _2  15:04:05 2006",
	"Mon 2006 January _2 15:04:05",

	"Mon Jan 2006 _2  15:04 UTC",
	"Mon Jan 2006 _2  15:04:05 UTC",
	"Mon Jan 2006 _2  15:04:05",
	"Mon 2006 Jan _2 15:04:05",

	"Jan 2006 _2  15:04 UTC",
	"Jan 2006 _2  15:04:05 UTC",
	"Jan 2006 _2  15:04:05",
	"2006 Jan _2 15:04:05",

	"_2 Jan 2006  15:04 UTC",
	"_2 Jan 2006  15:04:05 UTC",
	"_2 Jan 2006  15:04:05",
	"_2 2006 Jan  15:04:05",

	"Mon January 2006 _2  15:04 UTC",
	"Mon January 2006 _2  15:04:05 UTC",
	"Mon January 2006 _2  15:04:05",
	"Mon 2006 January _2 15:04:05",

	"January 2 2006 3:04 pm",
	"January 2 2006 3:04 pm",
	"January 2 2006 3:04 pm",
	"Jan 2 2006 03:04 pm",
	"Jan 2 2006 03:04 pm",
	"Jan 2 2006 3:04 pm",
	"Jan 2 06 3:04 pm",

	"January 2 2006 3:04:05 pm",
	"January 2 2006 3:04:05 pm",
	"January 2 2006 3:04:05 pm",
	"Jan 2 2006 03:04:05 pm",
	"Jan 2 2006 03:04:05 pm",
	"Jan 2 2006 3:04:05 pm",
	"Jan 2 06 3:04:05 pm",

	"Jan _2 03:04:05 pm 2006",
	"Jan _2 03:04:05 pm 2006",
	"Jan _2 3:04:05 pm 2006",
	"Jan _2 3:04:05 pm 2006",

	"January _2 03:04:05 pm 2006",
	"Jan _2 15:04:05 2006",
	"January _2 15:04:05 2006",
	"2006 January _2 15:04:05",
	"January 2006 _2 15:04:05",
	"2006 _2 January 15:04:05",
	"2006 _2 Jan 15:04:05",

	"Jan 2 15:04:05",
	"2 Jan 15:04:05",
	"15:04:05 2 Jan",

	"2006 15:04:05 2 Jan",
	"15:04:05 2006 2 Jan",
	"15:04:05 2 2006 Jan",

	"2 2006 Jan 15:04:05",
	"2006 Jan 2 15:04:05",

	"2006-01-02 3:04 pm",
	"2006-01-02 3:04pm",
	"2006-01-02 3:04 PM",
	"2006-01-02 3:04PM",
	"2006-01-02 15:04",

	"2006-01-02 3:04:05 pm",
	"2006-01-02 3:04:05pm",
	"2006-01-02 3:04:05 PM",
	"2006-01-02 3:04:05PM",
	"2006-01-02 15:04:05",

	"2006/01/02 3:04 pm",
	"2006/01/02 3:04pm",
	"2006/01/02 3:04 PM",
	"2006/01/02 3:04PM",
	"2006/01/02 15:04",

	"2006/01/02 3:04:05 pm",
	"2006/01/02 3:04:05pm",
	"2006/01/02 3:04:05 PM",
	"2006/01/02 3:04:05PM",
	"2006/01/02 15:04:05",

	"01/02/2006 3:04 pm",
	"01/02/2006 3:04pm",
	"01/02/2006 3:04 PM",
	"01/02/2006 3:04PM",
	"01/02/2006 15:04",

	"01/02/2006 3:04:05 pm",
	"01/02/2006 3:04:05pm",
	"01/02/2006 3:04:05 PM",
	"01/02/2006 3:04:05PM",
	"01/02/2006 15:04:05",

	"01.02.2006 3:04 pm",
	"01.02.2006 3:04pm",
	"01.02.2006 3:04 PM",
	"01.02.2006 3:04PM",
	"01.02.2006 15:04",

	"01.02.2006 3:04:05 pm",
	"01.02.2006 3:04:05pm",
	"01.02.2006 3:04:05 PM",
	"01.02.2006 3:04:05PM",
	"01.02.2006 15:04:05",

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

	"1.2.2006  15:04:05",
	"1/2/2006  15:04:05",
	"1_2_2006  15:04:05",
	"1 2 2006  15:04:05",

	"15:04 1.2.2006",
	"15:04 1/2/2006",
	"15:04 1_2_2006",
	"15:04 1 2 2006",

	"15:04:05 1.2.2006",
	"15:04:05 1/2/2006",
	"15:04:05 1_2_2006",
	"15:04:05 1 2 2006",

	"2006-01-02 15:04:05 +0000 UTC",
	"2006-01-02T15:04:05+07:00",
	"2006-01-02 15:04:05 +07:00",
	"2006-01-02T15:04:05-0700",

	"2006-01-02 15:04:05 -0700 MST",
	"2006-01-02 15:04:05 -0700 (MST)",
	"15:04:05 2006-01-02 -0700 MST",
	"15:04:05 2006-01-02 -0700 (MST)",

	"Mon Jan _2 15:04:05 MST 2006",
	"Mon Jan _2 15:04:05 -0700 2006",
	"02 Jan 06 15:04 MST",
	"02 Jan 06 15:04 -0700",
	"Monday, 02-Jan-06 15:04:05 MST",
	"Monday, 02-Jan-06 15:04:05 (MST)",
	"Mon _2 Jan 2006 15:04:05 MST",
	"Mon _2 Jan 2006 15:04:05 (MST)",
	"Mon _2 Jan 2006 15:04:05 -0700",
	"Mon _2 Jan 2006 15:04:05 -0700 MST",
	"Mon _2 Jan 2006 15:04:05 -0700 (MST)",
	"2006-01-02T15:04:05Z07:00",

	// Leave this last
	"2006-01-02T15:04:05.999999999Z07:00",
}

// getTime --
func (s DateTimeParse) GetTime() (time.Time, error) {

	t, err := ifEpoch(string(s))
	if err == nil {
		return t, err
	}

	st := strings.Join(strings.Fields(string(s)), " ")
	st = strings.ReplaceAll(st, ",", "")

	if t, err := time.Parse(lastused, st); err == nil {
		return t, err
	}

	for _, l := range layout {
		t, err := time.Parse(l, st)
		if err == nil {
			lastused = l
			return t, err
		}

	}

	return time.Time{}, errors.New("Time format is not in layout.")

}

func (s DateTimeParse) GetTimeInLocation(zone string) (time.Time, error) {

	loc, err := time.LoadLocation(zone)
	if err != nil {
		return time.Time{}, err
	}

	t, err := ifEpoch(string(s))
	if err == nil {
		zone_time := t.In(loc)
		return zone_time, err
	}

	st := strings.Join(strings.Fields(string(s)), " ")
	//fmt.Printf("-->%s\n", st)

	if t, err := time.ParseInLocation(lastused, st, loc); err == nil {
		return t, err
	}
	for _, l := range layout {
		t, err := time.ParseInLocation(l, st, loc)
		if err == nil {
			lastused = l
			return t, err
		}

	}

	return time.Time{}, errors.New("Time format is not in layout.")

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

// TimeIn gives time in zone
func (s DateTimeParse) TimeIn(zone string) (time.Time, error) {
	tt, err := DateTimeParse(s).GetTime()
	if err != nil {
		return tt, err
	}

	loc, err := time.LoadLocation(zone)
	if err != nil {
		return tt, err
	}

	zone_time := tt.In(loc)
	return zone_time, err

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
