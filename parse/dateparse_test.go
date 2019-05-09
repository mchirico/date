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

func TestDateTimeParse(t *testing.T) {
	s := " April 2, 2018, 6:45 pm"
	tt, err := DateTimeParse(s).GetTimeLoc()
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

func Test_ifEpoch(t *testing.T) {
	s := "1287621011948"

	tt, err := ifEpoch(s)
	if err != nil {
		t.Fatalf("Can't convert string")
	}

	if 1287621011 != tt.Unix() {
		t.Fatalf("Can't convert string")
	}

}

func Test_ifEpoch_Error(t *testing.T) {
	s := "1287jan1011948"

	_, err := ifEpoch(s)
	if err == nil {
		t.Fatalf("Invalid timestamp")
	}

}

func TestQuick(t *testing.T) {
	s := "Thu Mar 21 19:07:52 UTC 2019"
	tt, err := DateTimeParse(s).GetTimeLocSquish()
	fmt.Printf("_>%s<_ %v %v\n", s, tt, err)

}
