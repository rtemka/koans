package leetcode

import "strings"

// Given a pattern and a string s, find if s follows the same pattern.

// Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.

// Example 1:

// Input: pattern = "abba", s = "dog cat cat dog"
// Output: true
// Example 2:

// Input: pattern = "abba", s = "dog cat cat fish"
// Output: false
// Example 3:

// Input: pattern = "aaaa", s = "dog cat cat dog"
// Output: false

func wordPattern(pattern string, s string) bool {

	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}
	mp := make(map[byte]string, len(pattern))
	mw := make(map[string]byte, len(pattern))

	for i := range words {
		// если не нашли паттерн
		if _, ok := mp[pattern[i]]; !ok {
			mp[pattern[i]] = words[i]
			// и нет такого слова
			if _, ok := mw[words[i]]; !ok {
				mw[words[i]] = pattern[i]
			} else {
				// паттерна нет, а слово есть (
				return false
			}
		} else {
			// паттерн есть, но там не то слово
			if mp[pattern[i]] != words[i] {
				return false
			}
		}
	}

	return true
}
