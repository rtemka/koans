package leetcode

// Given two strings s and t, determine if they are isomorphic.

// Two strings s and t are isomorphic if the characters in s can be replaced to get t.

// All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.

// Example 1:

// Input: s = "egg", t = "add"
// Output: true
// Example 2:

// Input: s = "foo", t = "bar"
// Output: false
// Example 3:

// Input: s = "paper", t = "title"
// Output: true

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[byte]byte)
	m2 := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if v, ok := m[s[i]]; ok {
			if v != t[i] {
				return false
			}
		} else {
			m[s[i]] = t[i]
		}
		if v, ok := m2[t[i]]; ok {
			if v != s[i] {
				return false
			}
		} else {
			m2[t[i]] = s[i]
		}
	}

	return true
}
