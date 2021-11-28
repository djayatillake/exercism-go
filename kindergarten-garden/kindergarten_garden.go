package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

// Define the Garden type here.
type Garden []lookup

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

func NewGarden(diagram string, children []string) (*Garden, error) {
	if strings.Split(diagram, "\n")[0] != "" {
		return nil, errors.New("wrong format")
	}
	PlantMap := map[rune]string{
		'G': "grass", 'C': "clover", 'R': "radishes", 'V': "violets",
	}
	g := Garden{}
	sorted_children := append([]string{}, children...)
	sort.Strings(sorted_children)
	for i, child := range sorted_children {
		g = append(g, lookup{child: child, plants: []string{}, ok: true})
		for j, child2 := range sorted_children {
			if child == child2 && i != j {
				return nil, errors.New("dupe")
			}
		}
	}
	for _, v := range strings.Split(strings.TrimSpace(diagram), "\n") {
		len_v, len_c := len(v), len(children)
		if len_v/len_c != 2 || len_v%len_c != 0 {
			return nil, errors.New("lookup row wrong length")
		}
		for i := 0; i < len_v; i += 2 {
			v1, ok1 := PlantMap[rune(v[i])]
			v2, ok2 := PlantMap[rune(v[i+1])]
			if !ok1 || !ok2 {
				return nil, errors.New("invalid cup codes")
			}
			g[i/2].plants = append(g[i/2].plants, v1, v2)
		}
	}
	return &g, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	for _, lookup := range *g {
		if lookup.child == child {
			return lookup.plants, lookup.ok
		}
	}
	return nil, false
}
