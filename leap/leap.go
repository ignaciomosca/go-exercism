// Package leap implements IsLeapYear method to check if a year is leap year.
package leap

// IsLeapYear returns true if the year passed is leap year, false if it's not.
func IsLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	} else {
		return false
	}
}
