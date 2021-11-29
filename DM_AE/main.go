package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type change struct {
	id              string
	person          string
	event_type      string
	event_timestamp string
}

type snapshot struct {
	id      string
	started string
	closed  string
	people  string
}

type weekNumPeople struct {
	id             string
	week           string
	net_num_people int
	max_num_people int
}

const TimeForm = "2006-01-02 15:04:05 UTC"

func main() {
	reports := []weekNumPeople{}
	// iterating over each project
	for _, project := range current_snapshot {
		closed, err := time.Parse(TimeForm, project.closed)
		if err != nil {
			closed = time.Now()
		}
		closedWT := weekTruncUTC(closed)
		started, _ := time.Parse(TimeForm, project.started)
		startedWT := weekTruncUTC(started)
		// filtering changes for this project and sorting by time desc and remove before add
		changesFiltered := filterChanges(changes, project.id)
		sortByTimeAndEventType(changesFiltered, false, false)
		people := strings.Split(project.people, ",")
		// iterating from the closed or current week of project back to it's start
		for week := closedWT; week.After(startedWT.Add(time.Hour * 24 * -7)); week = week.Add(time.Hour * 24 * -7) {
			peopleStartNo := len(people)
			peopleMaxNos := []int{peopleStartNo}
			// initialising the report for the current project and week
			report := weekNumPeople{project.id, week.String()[0:10], peopleStartNo, peopleStartNo}
			// iterating over the filtered and sorted changes to update the report with the weeks changes
			for _, change := range changesFiltered {
				event_ts_time, _ := time.Parse(TimeForm, change.event_timestamp)
				next_week := week.Add(time.Hour * 24 * 7)
				/* in a dwh unnecessary scans of data can be dealt with by indexing the data,
				if built here this could allow only the needed data for the week to be scanned;
				possible improvement with more time. I've reduced some work by causing the loop
				to continue immediately for any event after the week and break at the first event
				before the week */
				if event_ts_time.After(next_week) || event_ts_time == next_week {
					continue
				} else if event_ts_time.Before(week) {
					break
				}
				if change.event_type == "removed" && !peopleContains(people, change.person) {
					people = append(people, change.person)
					report.net_num_people++
					peopleMaxNos = append(peopleMaxNos, report.net_num_people)
				} else if change.event_type == "added" && peopleContains(people, change.person) {
					people = remove(people, change.person)
					report.net_num_people--
				}
			}
			report.max_num_people = max(peopleMaxNos)
			reports = append(reports, report)
		}
	}
	// return reports
	fmt.Println("Output:")
	fmt.Println("id | week       | net_num_people | max_num_people")
	for _, report := range reports {
		fmt.Printf("%-3s| %s | %-14d | %-14d\n",
			report.id, report.week, report.net_num_people, report.max_num_people)
	}
	/* I am aware that this code generates a slightly different result to my sql solution
	this is due to this solution being unopinionated about which of consecutive same
	event types are valid; it will just use the first one that can be validly used depending
	on the state of the people in the project at that time. The sql solution would have
	not been able to handle processing the consecutive same event types and therefore
	they were filtered in an opinionated way favouring the earliest valid event. */
}

// weekTruncUTC truncates a time field to the start of the week it is in
func weekTruncUTC(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day()-int(t.Weekday()),
		0, 0, 0, 0, time.UTC)
}

// filterChanges filters []change by id
func filterChanges(changes []change, id string) (ret []change) {
	for _, change := range changes {
		if change.id == id {
			ret = append(ret, change)
		}
	}
	return
}

// remove returns a slice without the chosen item
func remove(people []string, person string) (ret []string) {
	for _, p := range people {
		if p != person {
			ret = append(ret, p)
		}
	}
	return
}

// sortByTimeAndEventType sorts changes by event timestamp and event type
//	in order as set by parameters
func sortByTimeAndEventType(changes []change, timeasc bool, addfirst bool) {
	sort.SliceStable(changes, func(i, j int) bool {
		ci, cj := changes[i], changes[j]
		switch {
		case ci.event_timestamp != cj.event_timestamp:
			if timeasc {
				return ci.event_timestamp < cj.event_timestamp
			} else {
				return ci.event_timestamp > cj.event_timestamp
			}
		default:
			if addfirst {
				return ci.event_type == "added"
			}
			return ci.event_type == "removed"
		}
	})
}

// peopleContains checks whether a person is in an array of people
func peopleContains(people []string, person string) bool {
	for _, p := range people {
		if p == person {
			return true
		}
	}
	return false
}

// max returns the max int from an []int
func max(nos []int) (ret int) {
	for _, no := range nos {
		if no > ret {
			ret = no
		}
	}
	return
}

/*Data - as the data was small, to avoid boilerplate code connecting to BQ and
unmarshalling etc I brought it in this way*/
var changes = []change{
	{
		id:              "1",
		person:          "p_e",
		event_type:      "removed",
		event_timestamp: "2021-04-12 00:00:00 UTC",
	},
	{
		id:              "1",
		person:          "p_d",
		event_type:      "added",
		event_timestamp: "2021-03-04 00:00:00 UTC",
	},
	{
		id:              "1",
		person:          "p_c",
		event_type:      "added",
		event_timestamp: "2021-01-27 00:00:00 UTC",
	},
	{
		id:              "2",
		person:          "p_b",
		event_type:      "removed",
		event_timestamp: "2021-09-23 00:00:00 UTC",
	},
	{
		id:              "2",
		person:          "p_b",
		event_type:      "added",
		event_timestamp: "2021-09-18 00:00:00 UTC",
	},
	{
		id:              "2",
		person:          "p_a",
		event_type:      "added",
		event_timestamp: "2021-08-26 00:00:00 UTC",
	},
	{
		id:              "2",
		person:          "p_a",
		event_type:      "removed",
		event_timestamp: "2021-08-01 00:00:00 UTC",
	},
	{
		id:              "2",
		person:          "p_b",
		event_type:      "added",
		event_timestamp: "2021-08-01 00:00:00 UTC",
	},
	{
		id:              "3",
		person:          "p_f",
		event_type:      "removed",
		event_timestamp: "2021-09-02 00:00:00 UTC",
	},
	{
		id:              "3",
		person:          "p_g",
		event_type:      "removed",
		event_timestamp: "2021-09-02 00:00:00 UTC",
	},
	{
		id:              "3",
		person:          "p_h",
		event_type:      "removed",
		event_timestamp: "2021-09-02 00:00:00 UTC",
	},
	{
		id:              "3",
		person:          "p_e",
		event_type:      "added",
		event_timestamp: "2021-08-27 00:00:00 UTC",
	},
	{
		id:              "3",
		person:          "p_e",
		event_type:      "removed",
		event_timestamp: "2021-08-27 00:00:00 UTC",
	},
	{
		id:              "3",
		person:          "p_a",
		event_type:      "added",
		event_timestamp: "2021-08-25 00:00:00 UTC",
	},
	{
		id:              "3",
		person:          "p_b",
		event_type:      "added",
		event_timestamp: "2021-08-25 00:00:00 UTC",
	},
}

var current_snapshot = []snapshot{
	{
		id:      "1",
		started: "2021-01-03 00:00:00 UTC",
		closed:  "2021-05-02 00:00:00 UTC",
		people:  "p_a,p_b,p_c,p_d",
	},
	{
		id:      "2",
		started: "2021-07-14 00:00:00 UTC",
		closed:  "",
		people:  "p_a,p_b,p_c,p_d",
	},
	{
		id:      "3",
		started: "2021-07-14 00:00:00 UTC",
		closed:  "2021-09-09 00:00:00 UTC",
		people:  "p_a,p_b,p_c,p_d",
	},
}
