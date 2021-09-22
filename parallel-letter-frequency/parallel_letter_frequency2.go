package letter

import (
	"sync"
	"unicode/utf8"
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

// SafeCounter is safe to use concurrently.
type SafeFreqMap struct {
	mu sync.Mutex
	v  FreqMap
}

// Inc increments the counter for the given key.
func (c *SafeFreqMap) Inc(key rune) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeFreqMap) Value(key rune) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func ConcurrentFrequency(s []string) FreqMap {
	m := SafeFreqMap{v: make(FreqMap)}
	var wg sync.WaitGroup
	for _, str := range s {
		wg.Add(utf8.RuneCountInString(str))
		for _, r := range str {
			go func(waitgroup *sync.WaitGroup, r rune) {
				m.Inc(r)
				waitgroup.Done()
			}(&wg, r)
		}
	}

	wg.Wait()

	return m.v
}
