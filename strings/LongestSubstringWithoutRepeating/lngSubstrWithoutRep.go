package leetcode

// #3
// https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	longest := 0
	m := make(map[byte]int)

	for lo, hi := -1, 0; hi < len(s); hi++ {
		if i, ok := m[s[hi]]; ok {
			lo = max(lo, i)
		}
		m[s[hi]] = hi
		longest = max(longest, hi-lo)
	}
	return longest
}
