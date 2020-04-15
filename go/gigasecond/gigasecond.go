// Package gigasecond implements function to add gigasecond to given time
package gigasecond

import (
	"time"
)

const gigasecond = time.Second * 1e9

// AddGigasecond adds gigasecond to a given time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond)
}
