package validpalindrome2

func validPalindrome(s string) bool {
	lo, hi := 0, len(s)-1
	for lo < hi {
		if s[lo] != s[hi] {
			return isPalindrome(s, lo+1, hi) || isPalindrome(s, lo, hi-1)
		}
		lo++
		hi--
	}
	return true
}

func isPalindrome(s string, lo, hi int) bool {
	for lo < hi {
		if s[lo] != s[hi] {
			return false
		}
		lo++
		hi--
	}
	return true
}
