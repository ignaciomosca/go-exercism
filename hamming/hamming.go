// Package hamming implements method Distance
// to calculate the Hamming distance of two differente DNA sequences.
package hamming

import (
	"errors"
	"unicode/utf8"
)

// Distance calculates the Hamming distance between two DNA sequences.
func Distance(a, b string) (int, error) {
	result := 0
	aBytes := []byte(a)
	bBytes := []byte(b)
	if utf8.RuneCount(aBytes) != utf8.RuneCount(bBytes) {
		return result, errors.New("sequences differ in lengths")
	}

	iB := 0
	for _, rA := range a {
		rB, rWidth := utf8.DecodeRuneInString(b[iB:])
		iB += rWidth

		if rA != rB {
			result++
		}
	}
	return result, nil
}
