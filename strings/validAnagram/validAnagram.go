// #242
// https://leetcode.com/problems/valid-anagram/description/
package vallidanagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var m ['z' + 1 - 'a']rune

	for _, r := range s {
		m[r-'a']++
	}
	for _, r := range t {
		m[r-'a']--
	}
	for _, r := range m {
		if r != 0 {
			return false
		}
	}
	return true
}

func isAnagramUnicode(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[rune]int, len(s))
	for _, r := range s {
		m[r]++
	}
	for _, r := range t {
		if _, ok := m[r]; !ok {
			return false
		} else {
			m[r]--
		}
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
