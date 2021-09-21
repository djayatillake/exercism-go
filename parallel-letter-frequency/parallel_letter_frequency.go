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

func conc_freq_thread(wg *sync.WaitGroup, s string, m FreqMap) {
	defer wg.Done()
	for _, r := range s {
		m[r]++
	}
}

func ConcurrentFrequency(s []string) FreqMap {
	var wg sync.WaitGroup
	var ms []FreqMap
	for _, str := range s {
		wg.Add(1)
		mi := FreqMap{}
		go conc_freq_thread(&wg, str, mi)
		ms = append(ms, mi)
	}

	wg.Wait()

	mf := FreqMap{}
	for _, m := range ms {
		for k, v := range m {
			mf[k] += v
		}
	}

	return mf
}
