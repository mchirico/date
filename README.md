[![Go](https://github.com/mchirico/date/actions/workflows/go.yml/badge.svg?branch=develop)](https://github.com/mchirico/date/actions/workflows/go.yml)

# date
Go Dateparse


## Install

```bash
go get -u github.com/mchirico/date/parse

```


### Usage

[playground](https://go.dev/play/p/K0rBJrvVjRn)
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

Works with timezones 

[playground](https://go.dev/play/p/6uPD1gedJNh)
```go
package main

import (
	"fmt"
	"github.com/mchirico/date/parse"
)

func main() {

	s := "1554934858234"
	tt, err := parse.DateTimeParse(s).TimeIn("America/New_York")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("tt: %v\n", tt)
	// tt: 2019-04-10 18:20:58.234 -0400 EDT

	tt, _ = parse.DateTimeParse(s).TimeIn("America/Chicago")
	fmt.Printf("tt: %v\n", tt)
	// tt: 2019-04-10 17:20:58.234 -0500 CDT

	tt, _ = parse.DateTimeParse(s).TimeIn("America/Detroit")
	fmt.Printf("tt: %v\n", tt)
	// tt: 2019-04-10 18:20:58.234 -0400 EDT

	tt, _ = parse.DateTimeParse(s).TimeIn("America/Denver")
	fmt.Printf("tt: %v\n", tt)
	// tt: 2019-04-10 16:20:58.234 -0600 MDT

	tt, _ = parse.DateTimeParse(s).TimeIn("America/Los_Angeles")
	fmt.Printf("tt: %v\n", tt)
	// tt: 2019-04-10 15:20:58.234 -0700 PDT

	tt, _ = parse.DateTimeParse(s).TimeIn("UTC")
	fmt.Printf("tt: %v\n", tt)
	// tt: 2019-04-10 22:20:58.234 +0000 UTC

	tt, _ = parse.DateTimeParse(s).TimeIn("Asia/Shanghai")
	fmt.Printf("tt: %v\n", tt)
	// tt: 2019-04-11 06:20:58.234 +0800 CST

}


```

Example of rounding down.

```go
package main

import (
	"fmt"
	"github.com/mchirico/date/parse"
	"time"
)

func main() {

	s := "Thu Mar 21 19:07:52 UTC 2019"
	tt, err := parse.DateTimeParse(s).TimeIn("America/New_York")

	fmt.Printf("tt: %s  err: %v\n", tt, err)
	// tt: 2019-03-21 15:07:52 -0400 EDT  err: <nil>

	fmt.Printf("tt: %s  err: %v\n", tt.Round(60*time.Minute), err)
	// tt: 2019-03-21 15:00:00 -0400 EDT  err: <nil>

}


```
