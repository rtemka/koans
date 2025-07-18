// #438
// https://leetcode.com/problems/find-all-anagrams-in-a-string/description
package leetcode

func findAnagrams(s string, p string) []int {
	var res []int
	// If len(p) exceeds len(s) then it impossible to form any anagram.
	if len(p) > len(s) {
		return res
	}
	// Setup 2 arrays to count frequencies of letters.
	expected, window := [26]int{}, [26]int{}
	// Populate 'expected' with the characters in string 'p'.
	for i := range p {
		expected[p[i]-'a']++
	}
	lo, hi := 0, 0
	for hi < len(s) {
		// Add the character at the right pointer
		// to 'window' before sliding the window.
		window[s[hi]-'a']++
		// If the window has reached the expected fixed length,
		// we advance the left pointer as well as the right pointer
		// to slide the window.
		if hi-lo+1 == len(p) {
			if window == expected {
				res = append(res, lo)
			}
			// Remove the character at the left pointer from
			// 'window' before advancing the left pointer.
			window[s[lo]-'a']--
			lo++
		}
		hi++
	}
	return res
}
