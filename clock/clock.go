// package clock provides a way to create a clock and
package clock

import "fmt"

// Clock implements a clock
type Clock struct {
	min int
}

const day_mins int = 1440

// New creates a clock, with min input and returns it
func New(hour, min int) Clock {
	min_t := (min+hour*60)%day_mins + day_mins
	return Clock{min_t % day_mins}
}

// Stringer method presents the time of a clock in "HH:MM" format
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.min/60, c.min%60)
}

// Add method adds minutes to the time of a clock
func (c Clock) Add(min int) Clock {
	return New(0, c.min+min)
}

// Subtract method adds minutes to the time of a clock
func (c Clock) Subtract(min int) Clock {
	return c.Add(-min)
}
