package letter

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
	mf := FreqMap{}
	c := make(chan FreqMap)
	for _, str := range s {
		go func(w string) {
			c <- Frequency(w)
		}(str)
	}

	for _ = range s {
		m := <-c
		for k, v := range m {
			mf[k] += v
		}
	}

	return mf
}
