package lib

import (
	"fmt"
	"strings"
	"time"
)

// returns time.Now()
//
// use along with finish() to time some code:
//
//		started := start()
//	 	// ...some code...
//	 	elapsed := finish(started)
func Start() time.Time {
	return time.Now()
}

// returns duration since prior start() or time.Time;
// use along with start() to time some code
func Finish(started time.Time) time.Duration {
	return time.Since(started)
}

// sanitizes stringified duration so the microsecond indicator
// can pass thru console pipes like find cmd;
// otherwise, it prints as "#┬╡s" instead of "#µs"
func SanitizeDuration(d time.Duration) string {
	s := fmt.Sprintf("%s", d)
	s = strings.Replace(s, "µs", "us", 1)
	return s
}

// converts nanos# to equivalent duration
func NanosToDuration[N int | int64 | float64](nanos N) time.Duration {
	return time.Duration(float64(nanos) * float64(time.Nanosecond))
}

// computes the average time from total / counter, converts to duration
func NanosAvgToDuration[N int | int64 | float64](totalNanos N, counter N) time.Duration {
	return NanosToDuration(float64(totalNanos) / float64(counter))
}
