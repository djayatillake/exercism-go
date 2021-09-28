// package tournament calculates scores and presents a league table
package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// type team stores the results of a team
type team struct {
	name           string
	MP, W, D, L, P int
}

// declaring league of teams
type league map[string]team

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
func result_opp(outcome string) (ret string, err error) {
	switch outcome {
	case "win":
		ret = "loss"
	case "loss":
		ret = "win"
	case "draw":
		ret = "draw"
	default:
		return "", errors.New("invalid outcome, needs to be win loss or draw")
	}
	return ret, nil
}

// league_entry checks for team in league and updates or inserts as relevant
func (l league) league_entry(t, o string) {
	if val, ok := l[t]; ok {
		val.update(o)
		l[t] = val
	} else {
		new := team{}
		new.update(o)
		l[t] = new
	}
}

// parse_results takes a result delimited by ; and inputs it into the league
func (l league) parse_result(result string) (err error) {
	result_sl := strings.Split(result, ";")
	if len(result_sl) != 3 {
		return errors.New("not enough values in result")
	}
	team_a := strings.TrimSpace(result_sl[0])
	team_b := strings.TrimSpace(result_sl[1])
	if team_a == team_b {
		return errors.New("teams in result are the same")
	}
	outcome_a := strings.TrimSpace(result_sl[2])
	outcome_b, err := result_opp(outcome_a)
	if err != nil {
		return err
	}
	l.league_entry(team_a, outcome_a)
	l.league_entry(team_b, outcome_b)
	return nil
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
func league_sort(l league) []team {
	tosort := []team{}
	for k, v := range l {
		t := team{name: k, MP: v.MP, W: v.W, D: v.D, L: v.L, P: v.P}
		tosort = append(tosort, t)
	}
	sortByPointsAndName(tosort)
	return tosort
}

// Print_results outputs header and results to an io.Writer
func Print_results(results []team, output io.Writer) {
	fmt.Fprintf(output, "Team                           | MP |  W |  D |  L |  P\n")
	for _, v := range results {
		fmt.Fprintf(output, "%-31s| %2d |  %d |  %d |  %d |  %d\n",
			v.name, v.MP, v.W, v.D, v.L, v.P)
	}
}

// Parse_results takes a set of results delimited by new lines and inputs it into the league and returns table
func Tally(input io.Reader, output io.Writer) error {
	scan := bufio.NewScanner(input)
	l := league{}
	for scan.Scan() {
		trimmed := strings.TrimSpace(scan.Text())
		if strings.HasPrefix(trimmed, "#") || trimmed == "" {
			continue
		}
		err := l.parse_result(trimmed)
		if err != nil {
			return err
		}
	}
	sorted_league := league_sort(l)
	Print_results(sorted_league, output)
	return nil
}
