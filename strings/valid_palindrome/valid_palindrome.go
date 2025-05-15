package validpalindrome

import (
	"bytes"
	"unicode"
)

func isPalindrome(s string) bool {
	str := toAsciiLower([]byte(s))
	for lo, hi := 0, len(str)-1; lo < hi; lo, hi = lo+1, hi-1 {
		if str[lo] != str[hi] {
			return false
		}
	}
	return true
}

func toAsciiLower(s []byte) []byte {
	return bytes.Map(func(r rune) rune {
		c := unicode.ToLower(r)
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			return c
		} else {
			return -1
		}
	}, s)
}
