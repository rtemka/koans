package strings

import "unicode/utf8"

// CountRunes count number of
// runes in provided string.
func CountRunes(s string) int {
	return countRunesManual(s)
}

// countRunesManual counts runes in
// string using regular for loop.
func countRunesManual(s string) int {
	var n int
	for range s {
		n++
	}
	return n
}

// countRunesBySTD counts runes in
// string using standard library utils.
func countRunesBySTD(s string) int {
	return utf8.RuneCountInString(s)
}
