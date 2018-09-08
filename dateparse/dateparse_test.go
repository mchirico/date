package dateparse

import (
	"fmt"
	"testing"
)

func TestDateTimeParse(t *testing.T) {
	s := " April 2, 2018, 6:45 pm"
	tt, err := DateTimeParse(s).getTimeLoc()
	if err != nil {
		t.Fatalf("Time gave error")
	}

	t3, _ := DateTimeParse(s).getTimeLocSquish()
	fmt.Printf("Time: %v \n",
		t3)

	s = " Apr 2, 2018, 6:45 am"
	tt, err = DateTimeParse(s).getTime()

	t2, err := DateTimeParse("2018-04-02 06:45:00 +0000 UTC").getTime()

	if tt != t2 {
		fmt.Printf("UTC not equal")
		t.Fatalf("UTC not equal %v,%v", tt, t2)
	}

	s = " Apr 2, 18, 6:45 am"
	tt, err = DateTimeParse(s).getTime()

	if tt != t2 {
		fmt.Printf("UTC not equal")
		t.Fatalf("UTC not equal %v,%v", tt, t2)
	}

	fmt.Println(tt.Unix())
}
