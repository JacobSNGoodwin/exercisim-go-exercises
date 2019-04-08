package letter

import (
	"fmt"
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	m   FreqMap
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key rune) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.m[key]++
	c.mux.Unlock()
}

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	fmt.Println(m)
	return m
}

// ConcurrentFrequency counts the frequency of each rune in an array of given texts and returns this
// data as a FreqMap
func ConcurrentFrequency(t []string) FreqMap {
	c := SafeCounter{m: make(map[rune]int)}
	for _, text := range t {
		go Frequency(text)
	}
	return c.m
}
