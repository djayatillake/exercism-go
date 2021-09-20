// package clock provides a way to create a clock and
package clock

import "fmt"

// Clock implements a clock
type Clock struct {
	min int
}

// New creates a clock, with hour and min inputs and outputs time as string
func New(hour, min int) Clock {
	min_t := (min+hour*60)%1440 + 1440
	return Clock{min_t % 1440}
}

// Stringer method presents the time of a clock in "HH:MM" format
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.min/60, c.min%60)
}

// Add method adds minutes to the time of a clock
func (c Clock) Add(min int) Clock {
	new_min := c.min + min
	return New(0, new_min)
}

// Subtract method adds minutes to the time of a clock
func (clock Clock) Subtract(min int) Clock {
	return clock.Add(-min)
}
