package main

import (
	"fmt"
	"time"
)

func main() {
	pl := fmt.Println

	now := time.Now()
	pl(now)

	then := time.Date(2018, 8, 17, 16, 32, 9, 7158496, time.UTC)
	pl(then)

	pl(then.Year())
	pl(then.Month())
	pl(then.Day())
	pl(then.Hour())
	pl(then.Minute())
	pl(then.Second())
	pl(then.Nanosecond())
	pl(then.Location())
	pl(then.Weekday())
	pl(then.Before(now))
	pl(then.After(now))
	pl(then.Equal(now))

	diff := now.Sub(then)
	pl(diff)
	pl(diff.Hours())
	pl(diff.Minutes())
	pl(diff.Seconds())
	pl(diff.Nanoseconds())

	pl(then.Add(diff))
	pl(then.Add(-diff))
}
