// package leap contains a function which check whether a year is a leap year
package leap

// IsLeapYear checks whether an int year is a leap year
func IsLeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}
