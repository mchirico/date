[![Build Status](https://travis-ci.org/mchirico/date.svg?branch=develop)](https://travis-ci.org/mchirico/date)

# date
Go Dateparse


## Install

```bash
go get -u github.com/mchirico/date/...

```


### Usage

```go

package main

import (
	"fmt"
	"github.com/mchirico/date/dateparse"
)

func main() {

	s := " April 2, 2018, 6:45 pm"
	tt, err := dateparse.DateTimeParse(s).GetTimeLoc()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("tt: %v\n", tt)
	// tt: 2018-04-02 14:45:00 -0400 EDT

}

```