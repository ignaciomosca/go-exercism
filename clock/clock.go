// Package clock exposes functions to implement a clock that handles times without dates.
package clock

import "fmt"

// Clock is an integer that represents the current time.
type Clock int

// New returns a new Clock with the formatted time in hours and minutes.
func New(hours, minutes int) Clock {
	minutesInAnHour := 1440
	c := (hours*60 + minutes) % (minutesInAnHour)
	if c < 0 {
		c += minutesInAnHour
	}
	return Clock(c)
}

// Add adds minutes to the clock
func (c Clock) Add(minutes int) Clock {
	return New(0, int(c)+minutes)
}

// Subtract substracts minutes to the clock.
func (c Clock) Subtract(minutes int) Clock {
	return New(0, int(c)-minutes)
}

// String shows the time in the HH:mm format.
func (c Clock) String() string {
	hours := int(c / 60)
	minutes := int(c % 60)
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}
