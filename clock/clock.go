// Package clock manages adding and subtracting times to a clock
package clock

import "fmt"

// BenchmarkAddMinutes-4           100000000               16.2 ns/op             0 B/op          0 allocs/op
// BenchmarkSubtractMinutes-4      100000000               16.2 ns/op             0 B/op          0 allocs/op
// BenchmarkCreateClocks-4         200000000                7.51 ns/op            0 B/op          0 allocs/op

// Clock stores the hours and minutes of a clock with minutes and hours rolled over
// I prefer this implementation to only using minues as it is clearer to me. Benchmarks were the same
// for using only minutes in type clock
type Clock struct {
	m int
}

// New returns a clock of type Clock which represents hours and minutes.
// It accounts for overflow in hours and minutes along with negative hours
// and minutes
func New(h int, m int) Clock {
	minutes := (60*h + m) % 1440
	// 1440 minutes in a day roll over
	if minutes < 0 {
		minutes += 1440
	}

	return Clock{m: minutes}
}

// Add adds a minutes to a clock c of type Clock and returns a new Clock
func (c Clock) Add(a int) Clock {
	return New(c.m/60, c.m%60+a)
}

// Subtract subtracts a minutes to a clock c of type Clock and returns a new Clock
func (c Clock) Subtract(a int) Clock {
	return New(c.m/60, c.m%60-a)
}

// String converts the clock to a string of format hh:mm
func (c *Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.m/60, c.m%60)
}
