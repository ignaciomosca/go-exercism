// Package luhn exposes functions to determine if a string is valid
package luhn

import "strings"

// Valid determines if a string is valid according to the Luhn's algorithm.
func Valid(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	result := 0
	if len(s) <= 1 {
		return false
	}
	oddLength := len(s) % 2
	for i, digit := range s {
		number := int(digit - '0')
		if number >= 0 && number <= 9 {
			if i%2 == oddLength {
				number *= 2
				if number > 9 {
					number -= 9
				}
			}
			result += number
		} else { // if it's not a number return false
			return false
		}
	}
	return result%10 == 0
}
