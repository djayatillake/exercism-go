package booking

import (
	"time"
	"fmt"
)

const Form = "1/2/2006 15:04:05"

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {	
	t, _ := time.Parse(Form, date)
	return t
}

const Form2 = "January 2, 2006 15:04:05"

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	t, _ := time.Parse(Form2, date)
	return t.Before(time.Now())
}

const Form3 = "Monday, January 2, 2006 15:04:05"

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	t, _ := time.Parse(Form3, date)
	return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	t := Schedule(date)
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", 
							t.Weekday(), t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
