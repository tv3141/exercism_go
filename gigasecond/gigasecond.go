// Package gigasecond calculates the moment when someone has lived for 109 seconds.
package gigasecond

import (
    "time"
    "math"
)

// AddGigasecond adds 10^9 seconds to a timestamp.
func AddGigasecond(t time.Time) time.Time {
    return t.Add(time.Duration(math.Pow(10, 9)).Second)
}
