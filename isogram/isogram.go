// Package isogram exposes the IsIsogram function
package isogram

import "strings"

// IsIsogram returns true when a word has no repeated letters (hyphens and spaces can be repeated)
func IsIsogram(word string) bool {
	presentLetters := make(map[rune]bool)
	for _, w := range strings.ToUpper(word) {
		if w == '-' || w == ' ' {
			continue
		}
		if presentLetters[w] {
			return false
		}
		presentLetters[w] = true
	}
	return true
}
