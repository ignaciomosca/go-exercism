// Package raindrops exposes function Convert
package raindrops

import (
	"strconv"
)

// Convert makes a conversion from int to string based on the numbers factors
func Convert(input int) string {
	ret := ""
	if input%3 == 0 {
		ret += "Pling"
	}
	if input%5 == 0 {
		ret += "Plang"
	}
	if input%7 == 0 {
		ret += "Plong"
	}
	if ret == "" {
		return strconv.Itoa(input)
	}
	return ret
}
