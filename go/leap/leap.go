// Package leap implements function that determines if a given year is a leap year or not
package leap

// IsLeapYear returns true if given year is a leap year
func IsLeapYear(y int) bool {
	if y%4 != 0 || (y%100 == 0 && y%400 != 0) {
		return false
	}
	return true
}
