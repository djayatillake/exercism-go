// package tournament calculates scores and presents a league table
package tournament

import (
	"fmt"
	"sort"
	"strings"
)

// type team stores the results of a team
type team struct {
	name           string
	MP, W, D, L, P int
}

// declaring league of teams
var league = map[string]team{}

// calc method, calculates MP and P from W D L
func (t *team) calc() {
	t.MP = t.W + t.D + t.L
	t.P = t.W*3 + t.D*1
}

// update method, updates W D L and runs calc()
func (t *team) update(outcome string) {
	switch outcome {
	case "win":
		t.W++
	case "draw":
		t.D++
	case "loss":
		t.L++
	}
	t.calc()
}

// result_opp for other team
func result_opp(outcome string) (ret string) {
	switch outcome {
	case "win":
		ret = "loss"
	case "loss":
		ret = "win"
	case "draw":
		ret = "draw"
	}
	return
}

// league_entry checks for team in league and updates or inserts as relevant
func league_entry(t, o string) {
	if val, ok := league[t]; ok {
		val.update(o)
		league[t] = val
	} else {
		new := team{}
		new.update(o)
		league[t] = new
	}
}

// parse_results takes a result delimited by ; and inputs it into the league
func parse_result(result string) {
	result_sl := strings.Split(result, ";")
	team_a := strings.TrimSpace(result_sl[0])
	team_b := strings.TrimSpace(result_sl[1])
	outcome_a := strings.TrimSpace(result_sl[2])
	outcome_b := result_opp(outcome_a)

	league_entry(team_a, outcome_a)
	league_entry(team_b, outcome_b)
}

func sortByPointsAndName(teams []team) {
	sort.SliceStable(teams, func(i, j int) bool {
		ti, tj := teams[i], teams[j]
		switch {
		case ti.P != tj.P:
			return ti.P > tj.P
		default:
			return ti.name < tj.name
		}
	})
}

// league sort accepts a league dictionary and returns a slice of teams sorted by points
func league_sort(l map[string]team) []team {
	tosort := []team{}
	for k, v := range l {
		t := team{name: k, MP: v.MP, W: v.W, D: v.D, L: v.L, P: v.P}
		tosort = append(tosort, t)
	}
	sortByPointsAndName(tosort)
	return tosort
}

func Print_results(results []team) (ret string) {
	ret += "\nTeam                           | MP |  W |  D |  L |  P\n"
	for _, v := range results {
		ret += fmt.Sprintf("%-31s| %2d |  %d |  %d |  %d |  %d\n",
			v.name, v.MP, v.W, v.D, v.L, v.P)
	}
	return
}

// Parse_results takes a set of results delimited by new lines and inputs it into the league and returns table
func Tally(results string) string {
	for _, v := range strings.Split(strings.TrimSpace(results), "\n") {
		parse_result(v)
	}
	sorted_league := league_sort(league)
	return Print_results(sorted_league)
}
