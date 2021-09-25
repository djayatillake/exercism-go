package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Robot struct {
	name string
}

var (
	namePool = genNames()
	nameIdx  = 0
)

const maxIdx = 676_000

func genNames() []string {
	pos := 0
	var w strings.Builder
	pool := make([]string, maxIdx)
	for c := 'A'; c < '['; c++ {
		for c2 := 'A'; c2 < '['; c2++ {
			for num := 0; num < 1000; num++ {
				w.WriteRune(c)
				w.WriteRune(c2)
				w.WriteString(fmt.Sprintf("%03d", num))
				pool[pos] = w.String()
				w.Reset()
				pos++
			}
		}
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(pool), func(i, j int) { pool[i], pool[j] = pool[j], pool[i] })
	return pool
}

func getName() (string, error) {
	if nameIdx == maxIdx {
		return "", errors.New("i came to kick ass and take names and I'm all outta names")
	}
	name := namePool[nameIdx]
	nameIdx++
	return name, nil
}

func (i *Robot) Name() (string, error) {
	if i.name == "" {
		err := i.Reset()
		if err != nil {
			return "", err
		}
	}
	return i.name, nil
}

func (i *Robot) Reset() error {
	name, err := getName()
	if err != nil {
		return err
	}
	i.name = name
	return nil
}
