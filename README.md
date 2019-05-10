[![Build Status](https://travis-ci.org/mchirico/date.svg?branch=develop)](https://travis-ci.org/mchirico/date)
[![codecov](https://codecov.io/gh/mchirico/date/branch/develop/graph/badge.svg)](https://codecov.io/gh/mchirico/date)

# date
Go Dateparse


## Install

```bash
go get -u github.com/mchirico/date/parse

```


### Usage

```go

package main

import (
	"fmt"
	"github.com/mchirico/date/parse"
)

func main() {

	s := "Sep  8  13:24:18 "
	tt, err := parse.DateTimeParse(s).NewYork()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("tt: %v\n", tt)
	// tt: 2018-09-08 13:24:18 -0400 EDT

}

```

Here's another example:

```go
package main

import (
	"fmt"
	"github.com/mchirico/date/parse"
)

func main() {

	s := "1554934858234"
	tt, err := parse.DateTimeParse(s).NewYork()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("tt: %v\n", tt)
	// tt: 2019-04-10 22:20:58.234 -0400 EDT

}


```

