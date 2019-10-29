// Package grains provides functions to calculate the number of grains
// a king had to pay a peaseant that saved his life.
package grains

import "errors"

// Square calculates the value of 2^(n-1) of a given number n.
func Square(i int) (uint64, error) {
	if i <= 0 || i > 64 {
		return 0, errors.New("math: out of range value")
	}
	return (1 << (uint64(i) - 1)), nil
}

// Total calculates the sum of squares from 1 to 64.
func Total() uint64 {
	return 1<<64 - 1
}
