package letter

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(s []string) FreqMap {
	var wg sync.WaitGroup
	mf := FreqMap{}
	wg.Add(len(s))
	c := make(chan FreqMap, 3)

	for _, str := range s {
		go func(waitgroup *sync.WaitGroup, w string) {
			c <- Frequency(w)
			waitgroup.Done()
		}(&wg, str)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for m := range c {
		for k, v := range m {
			mf[k] += v
		}
	}

	return mf
}
