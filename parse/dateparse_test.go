package parse

import (
	"fmt"
	"testing"
	"time"
)

func TestLoctoUTC(t *testing.T) {
	s := "Sep  8  13:24:18 "
	tt, err := DateTimeParse(s).NewYork()
	if err != nil {
		t.Fatalf("Time gave error")
	}

	if tt.Year() != time.Now().Year() {
		t.Fatalf("Should default to current year: %v %v\n",
			tt, time.Now().Year())
	}

	year := fmt.Sprintf("%d-09-08 17:24:18 +0000 UTC", tt.Year())

	t2, err := DateTimeParse(year).GetTime()

	loc, err := time.LoadLocation("America/New_York")

	if tt.In(loc) != t2.In(loc) {
		t.Fatalf("Times should be equal: %v,%v", tt, t2)
	}

	fmt.Println(tt)

}

func TestErrorChecking(t *testing.T) {
	s := " Aprill 2, 2018, 6:45 pm"
	_, err := DateTimeParse(s).GetTimeLoc()
	if err == nil {
		t.Fatalf("Should be error... two ll's in Aprill")
	}

	r, err := DateTimeParse(s).GetTimeLocSquish()
	if err == nil {
		t.Fatalf("Should be error... two ll's in Aprill")
	}

	if r != "" {
		t.Fatalf("Result should be empty string")
	}

	r, err = DateTimeParse(s).GetTimeLocHRminS()

	if err == nil {
		t.Fatalf("Should be error... two ll's in Aprill")
	}
	if r != "" {
		t.Fatalf("Result should be empty string")
	}

	good := " April 2, 2018, 6:45 pm"
	_, err = DateTimeParse(good).DaysBetween(s)
	if err == nil {
		t.Fatalf("Should be error... two ll's in Aprill")
	}

	_, err = DateTimeParse(s).DaysBetween(good)
	if err == nil {
		t.Fatalf("Should be error... two ll's in Aprill")
	}

	_, err = DateTimeParse(good).DaysFrom(s)
	if err == nil {
		t.Fatalf("Should be error... two ll's in Aprill")
	}

	_, err = DateTimeParse(s).DaysFrom(good)
	if err == nil {
		t.Fatalf("Should be error... two ll's in Aprill")
	}

}

func TestDateTimeParse(t *testing.T) {
	s := " April 2, 2018, 6:45 pm"
	tt, err := DateTimeParse(s).GetTimeLoc()
	if err != nil {
		t.Fatalf("Time gave error")
	}

	s = " Wed, 9 Sep 2020 11:34:00 -0400 (EDT)"
	tt, err = DateTimeParse(s).GetTimeLoc()
	if err != nil {
		t.Fatalf("Time gave error")
	}

	t3, _ := DateTimeParse(s).GetTimeLocSquish()
	fmt.Printf("Time: %v \n",
		t3)

	s = " Apr 2, 2018, 6:45 am"
	tt, err = DateTimeParse(s).GetTime()

	t2, err := DateTimeParse("2018-04-02 06:45:00 +0000 UTC").GetTime()

	if tt != t2 {
		fmt.Printf("UTC not equal")
		t.Fatalf("UTC not equal %v,%v", tt, t2)
	}

	s = " Apr 2, 18, 6:45 am"
	tt, err = DateTimeParse(s).GetTime()

	if tt != t2 {
		fmt.Printf("UTC not equal")
		t.Fatalf("UTC not equal %v,%v", tt, t2)
	}

	fmt.Println(tt.Unix())
}

func TestDateTimeParse_DaysFrom(t *testing.T) {
	s := " April 2, 2018, 6:45 pm"
	s2 := " April 5, 2018, 6:45 pm"

	days, err := DateTimeParse(s).DaysFrom(s2)
	if err != nil {
		t.Fatalf("Time gave error")
	}

	if days != 3 {
		t.Fatalf("Expected 3, Got: %d\n", days)
	}
}

func TestDateTimeParse_DaysBetween(t *testing.T) {
	s2 := " April 2, 2018, 6:45 pm"
	s := " April 5, 2018, 6:45 pm"

	days, err := DateTimeParse(s).DaysBetween(s2)
	if err != nil {
		t.Fatalf("Time gave error")
	}

	if days != 3 {
		t.Fatalf("Expected 3, Got: %d\n", days)
	}

}

func TestDateTimeParse_Long(t *testing.T) {
	r, err := DateTimeParse("Thu Mar 21 18:54:16 UTC 2019").GetTimeLocSquish()

	if err != nil {
		t.FailNow()
	}

	if r != "18:50" {
		t.Fatalf("Expected: 18:50, Got: %s\n", r)
	}

}

func TestDateTimeParse_NewYork(t *testing.T) {
	r, err := DateTimeParse("Thu Mar 21 18:54:16 UTC 2019").NewYork()

	if err != nil {
		t.FailNow()
	}

	expected := "2019-03-21 18:54:16 -0400 EDT"
	if r.String() != expected {
		t.Fatalf("Expected: %s, Got: %s", expected, r.String())
	}

	r, err = DateTimeParse("aThu Mar 21 18:54:16 UTC 2019").NewYork()
	if err == nil {
		t.FailNow()
	}

}

func TestDateTimeParse_GetTimeLocHRminS(t *testing.T) {
	s := "1287621011948"
	r, err := DateTimeParse(s).GetTimeLocHRminS()

	t.Logf("r: %v\n", r)
	if err != nil {
		t.FailNow()
	}

	if r != "20:30" {
		// Testing on travis, which is UTC
		if r != "00:30" {
			t.Fatalf("Expected: 20:30, got: %v\n", r)
		}
	}

}

func Test_ifEpoch(t *testing.T) {
	s := "1287621011948"

	tt, err := ifEpoch(s)
	if err != nil {
		t.Fatalf("Can't convert string")
	}

	if 1287621011 != tt.Unix() {
		t.Fatalf("Can't convert string")
	}
	if tt.Nanosecond() != 948000000 {
		t.Fatalf("Expected: %d\n", 948000000)
	}

	s = "1287621011"

	tt, err = ifEpoch(s)
	if err != nil {
		t.Fatalf("Can't convert string")
	}

	if tt.Unix() != 1287621011 {
		t.Fatalf("tt.Unix() failed.")
	}

}

func Test_ifEpoch_Error(t *testing.T) {
	s := "1287jan1011948"

	_, err := ifEpoch(s)
	if err == nil {
		t.Fatalf("Invalid timestamp")
	}

}

func Test_EpochFull(t *testing.T) {
	s := "1287621011948"
	r, err := DateTimeParse(s).NewYork()
	if err != nil {
		t.FailNow()
	}

	if r.Nanosecond() != 948000000 {
		t.FailNow()
	}

	if r.String() != "2010-10-21 00:30:11.948 -0400 EDT" {
		t.FailNow()
	}

}

func TestQuick(t *testing.T) {
	s := "Thu Mar 21 19:07:52 UTC 2019"
	tt, err := DateTimeParse(s).GetTimeLocSquish()
	fmt.Printf("_>%s<_ %v %v\n", s, tt, err)

}

func TestTimeIn(t *testing.T) {

	s := "Thu Mar 21 19:07:52 UTC 2019"
	tt, err := DateTimeParse(s).TimeIn("America/New_York")
	if err != nil {
		t.Fatalf("Should not return error: %v\n", err)
	}
	if tt.String() != "2019-03-21 15:07:52 -0400 EDT" {
		t.FailNow()
	}

	tt, err = DateTimeParse(s).TimeIn("Invalid America/New_York")
	if err == nil {
		t.FailNow()
	}

	tt, err = DateTimeParse("junk").TimeIn("America/New_York")
	if err == nil {
		t.FailNow()
	}

	tt, err = DateTimeParse("Wed, 9 Sep 2020 09:06:34 -0700").TimeIn("America/New_York")
	if err != nil {
		t.FailNow()
	}

}

func TestRound(t *testing.T) {

	s := "Thu Mar 21 19:07:52 UTC 2019"
	tt, err := DateTimeParse(s).TimeIn("America/New_York")

	t.Logf("tt: %s  err: %v\n", tt, err)
	t.Logf("tt: %s  err: %v\n", tt.Round(60*time.Minute), err)

	expected, err := DateTimeParse("15:00:00 2019-03-21   -0400 EDT").GetTime()

	if !tt.Round(60 * time.Minute).Equal(expected) {
		t.FailNow()
	}
	t.Logf("tt: %s  expected: %v\n", tt.Round(60*time.Minute), expected)

}

func TestDateTimeParse_GetTimeInLocation(t *testing.T) {

	s := "16:02:23 5.12.2019"
	tt, err := DateTimeParse(s).GetTimeInLocation("Asia/Shanghai")
	if err != nil {
		t.FailNow()
	}
	if tt.String() != "2019-05-12 16:02:23 +0800 CST" {
		t.FailNow()
	}

	tt, err = DateTimeParse(s).GetTimeInLocation("bogus")
	if err == nil {
		t.Logf("tt: %s  err: %v\n", tt, err)
		t.FailNow()
	}

	tt, err = DateTimeParse("123456789ten").GetTimeInLocation("bogus")
	if err == nil {
		t.Logf("tt: %s  err: %v\n", tt, err)
		t.FailNow()
	}

	tt, err = DateTimeParse("1287621011948").GetTimeInLocation("America/New_York")
	if err != nil {
		t.Logf("tt: %s  err: %v\n", tt, err)
		t.FailNow()
	}
	if tt.String() != "2010-10-20 20:30:11.948 -0400 EDT" {
		t.Logf("tt: %v  err: %v\n", tt.String(), err)
		t.FailNow()
	}

	tt, err = DateTimeParse("2006-01-02T15:04:05.999999999Z07:00").GetTimeInLocation("UTC")
	if err == nil {
		t.Logf("tt: %s  err: %v\n", tt, err)
		t.FailNow()
	}

}

func TestEdge(t *testing.T) {
	s := "16:02:23 5.12.2019"
	tt, err := DateTimeParse(s).NewYork()
	t.Logf("tt: %s  err: %v\n", tt, err)

	s2 := "Sun May 12 16:02:23  EDT 2019"
	t2, err := DateTimeParse(s2).GetTimeInLocation("America/New_York")

	t.Logf("t2: %s  err: %v\n", t2, err)

	if !t2.Equal(tt) {
		t.FailNow()
	}

	// ParseInLocation(layout, value string, loc *Location) (Time, error)

}
