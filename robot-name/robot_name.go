// Package robotname contains functions to generate a new random name for a robot
package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.NewSource(time.Now().UnixNano())
}

// 1000 possible numbers times 26² (two letters from A to Z)
const maxNames = 26 * 26 * 10 * 10 * 10

// Robot is a struct that represents a robot, which only has one attribute, its name
type Robot struct {
	name string
}

var used = make(map[string]bool)

// Name returns a new name for the Robot
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	if len(used) == maxNames {
		return "", errors.New("no more unique names")
	}
	r.name = generateNewName()
	for used[r.name] {
		r.name = generateNewName()
	}
	used[r.name] = true
	return r.name, nil
}

// Reset resets the robots name which has to have the following format CCNNN
// Where C's are characters and N are numbers
func (r *Robot) Reset() {
	r.name = ""
}

func generateNewName() string {
	r1 := string(rand.Intn(26) + 'A')
	r2 := string(rand.Intn(26) + 'A')
	num := rand.Intn(1000)
	return fmt.Sprintf("%s%s%03d", r1, r2, num)
}
