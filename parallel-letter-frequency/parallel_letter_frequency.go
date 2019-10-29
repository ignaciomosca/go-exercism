// Package letter implements functions to count words.
package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// ConcurrentFrequency counts the frequency of each rune for a given strings array and returns a FreqMap.
func ConcurrentFrequency(strings []string) FreqMap {
	m := FreqMap{}
	c := make(chan FreqMap, 10)
	for _, v := range strings {
		go func(s string) {
			c <- Frequency(s)
		}(v)
	}
	for range strings {
		for r, frequency := range <-c {
			m[r] += frequency
		}
	}
	return m
}

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}
