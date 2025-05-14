// #541
// https://leetcode.com/problems/reverse-string-ii/description/
package reversestring2

func reverseStr(s string, k int) string {
	sb := []byte(s)
	for i := 0; i < len(sb); i = i + 2*k {
		if i+k < len(sb) {
			rev(sb[i : i+k])
		} else {
			rev(sb[i:])
		}

	}
	return string(sb)
}

func rev(s []byte) {
	lo, hi := 0, len(s)-1
	for lo < hi {
		s[lo], s[hi] = s[hi], s[lo]
		lo, hi = lo+1, hi-1
	}
}
