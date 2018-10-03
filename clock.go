package timesrs

import "time"

// Clock provides the time of the operation.
type Clock interface {
	Time() time.Time
}

type utcClock struct{}

// Time returns the current UTC time (zero offset).
func (*utcClock) Time() time.Time {
	return time.Now().UTC()
}

// LocalClock provides the current local time.
type LocalClock struct{}

// Time returns the current local time.
func (*LocalClock) Time() time.Time {
	return time.Now()
}

// DefaultClock is the default implementation of the clock.
var DefaultClock utcClock
