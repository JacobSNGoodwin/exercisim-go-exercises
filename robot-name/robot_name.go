package robotname

import (
	"math/rand"
)

// Robot blabla
type Robot struct {
	name string
}

// Name sets the name property of a Robot to a specific random value
func (r *Robot) Name() (string, error) {
	// a new robot off the factory floor getting its name
	if r.name == "" {
		r.name = createRobotName()
	}
	// if it already has a name, just return it
	return r.name, nil
}

// Reset again give the robot a new name
func (r *Robot) Reset() {
	r.name = createRobotName()
}

// createRobotName creates a robot string of the following format ^[A-Z]{2}\d{3}$
func createRobotName() string {
	// min and max codepoints of capital letters
	minl, maxl := int32(65), int32(90)
	// min and max codepoints of digits
	minn, maxn := int32(48), int32(57)

	rs := []rune{
		rune(rand.Int31n(maxl-minl) + minl),
		rune(rand.Int31n(maxl-minl) + minl),
		rune(rand.Int31n(maxn-minn) + minn),
		rune(rand.Int31n(maxn-minn) + minn),
		rune(rand.Int31n(maxn-minn) + minn),
	}

	return string(rs)
}
