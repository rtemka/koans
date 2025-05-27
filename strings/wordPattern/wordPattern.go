package leetcode

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}
	patternToWord := make(map[byte]string, len(words))
	wordToPattern := make(map[string]byte, len(words))
	for i := range len(words) {
		char := pattern[i]
		word := words[i]
		if v, ok := patternToWord[char]; ok && v != word {
			return false
		}
		if v, ok := wordToPattern[word]; ok && v != char {
			return false
		}
		patternToWord[char] = word
		wordToPattern[word] = char
	}
	return true
}

func WordPattern2022(pattern string, s string) bool {

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
