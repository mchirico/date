package dateparse

import (
	"fmt"
	"testing"
	"time"
)

func TestLoctoUTC(t *testing.T) {
	s := "Sep  8 13:24:18 "
	tt, err := DateTimeParse(s).LoctoUTC()
	if err != nil {
		t.Fatalf("Time gave error")
	}

	if tt.Year() != time.Now().Year() {
		t.Fatalf("Should default to current year: %v %v\n",
			tt, time.Now().Year())
	}

	t2, err := DateTimeParse("2018-09-08 17:24:18 +0000 UTC").GetTime()

	if tt != t2 {
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
