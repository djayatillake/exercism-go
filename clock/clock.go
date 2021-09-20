// package clock provides a way to create a clock and
package clock

import "fmt"

// Clock implements a clock
type Clock struct {
	hour int
	min  int
}

// New creates a clock, with hour and min inputs and outputs time as string
func New(hour, min int) Clock {
	min_r := min%1440 + 1440
	hour_r := hour%24 + 24
	clock := Clock{hour: (hour_r + min_r/60) % 24, min: min_r % 60}
	return clock
}

// dig_string helper func converts hour or mins to two digit string
func dig_string(dig int) string {
	if dig < 10 {
		return "0" + fmt.Sprint(dig)
	} else {
		return fmt.Sprint(dig)
	}
}

// clocktime method presents the time of a clock in "HH:MM" format
func (clock Clock) String() string {
	return dig_string(clock.hour) + ":" + dig_string(clock.min)
}

// Add method adds minutes to the time of a clock
func (clock Clock) Add(min int) Clock {
	min_r := min%1440 + 1440
	clock.hour = (clock.hour + (clock.min+min_r)/60) % 24
	clock.min = (clock.min + min_r) % 60
	return clock
}

// Subtract method adds minutes to the time of a clock
func (clock Clock) Subtract(min int) Clock {
	return clock.Add(-min)
}
