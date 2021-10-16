package meetup

import (
	"fmt"
	"time"
)

type WeekSchedule string

var (
	Teenth WeekSchedule = "teenth"
	First  WeekSchedule = "first"
	Second WeekSchedule = "second"
	Third  WeekSchedule = "third"
	Fourth WeekSchedule = "fourth"
	Last   WeekSchedule = "last"
)

const Form = "January 2 2006"

func Day(week WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	days := []int{}
	for i := 1; i < 32; i++ {
		date := fmt.Sprintf("%s %d %d", month.String(), i, year)
		t, err := time.Parse(Form, date)
		if err == nil && t.Weekday() == weekday {
			days = append(days, i)
		}
	}
	switch week {
	case First:
		return days[0]
	case Second:
		return days[1]
	case Third:
		return days[2]
	case Fourth:
		return days[3]
	case Last:
		return days[len(days)-1]
	default:
		if days[1] < 13 {
			return days[2]
		}
		return days[1]
	}
}
