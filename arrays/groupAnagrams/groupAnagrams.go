package leetcode

import (
	"sort"
)

// Given an array of strings strs, group the anagrams together. You can return the answer in any order.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
// typically using all the original letters exactly once.

// Example 1:

// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

// Example 2:

// Input: strs = [""]
// Output: [[""]]

// Example 3:

// Input: strs = ["a"]
// Output: [["a"]]

// Constraints:

// 1 <= strs.length <= 104
// 0 <= strs[i].length <= 100
// strs[i] consists of lowercase English letters.

func groupAnagrams(strs []string) [][]string {

	var res [][]string
	var already = make([]bool, len(strs))

	for i := range strs {
		if already[i] {
			continue
		}

		var ms = []string{strs[i]}
		trs := sortRunes([]rune(strs[i])) // target runes sorted

		for j := i + 1; j < len(strs); j++ {
			if already[j] || len(strs[i]) != len(strs[j]) {
				continue
			}
			ers := sortRunes([]rune(strs[j])) // examined runes sorted
			if compareRuneSlice(trs, ers) {
				already[j] = true
				ms = append(ms, strs[j])
			}
		}

		res = append(res, ms)
	}

	return res
}

func compareRuneSlice(left, right []rune) bool {
	if len(left) != len(right) {
		return false
	}
	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}

func sortRunes(rs []rune) []rune {
	sort.Slice(rs, func(i, j int) bool {
		return rs[i] < rs[j]
	})
	return rs
}

func groupAnagramsMap(strs []string) [][]string {

	var res [][]string
	var already = make([]bool, len(strs))

	for i := range strs {
		if already[i] {
			continue
		}

		var ms = []string{strs[i]}
		var tm = make(map[rune]int) // target map
		for _, r := range strs[i] {
			tm[r]++
		}

		for j := i + 1; j < len(strs); j++ {
			if already[j] || len(strs[i]) != len(strs[j]) {
				continue
			}
			var em = make(map[rune]int) // examined map
			for _, r := range strs[j] {
				em[r]++
			}

			if equalMaps(tm, em) {
				already[j] = true
				ms = append(ms, strs[j])
			}
		}

		res = append(res, ms)
	}

	return res
}

func equalMaps(x, y map[rune]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, v := range x {
		if i, ok := y[k]; !ok || v != i {
			return false
		}
	}

	return true
}

func groupAnagramsFast(strs []string) [][]string {
	m := make(map[int][]string, 1)
	for i := 0; i < len(strs); i++ {
		m[hash(strs[i])] = append(m[hash(strs[i])], strs[i])
	}
	result := make([][]string, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}

	return result
}

func hash(s string) int {
	multiplication := 1
	sum := 0
	for i := 0; i < len(s); i++ {
		multiplication *= int(s[i])
		sum += int(s[i])
	}

	return multiplication + sum
}
