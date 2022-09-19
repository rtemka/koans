package leetcode

import "log"

// You are given a string s and an array of strings words. All the strings of words are of the same length.

// A concatenated substring in s is a substring that contains all the strings of any permutation of words concatenated.

// For example, if words = ["ab","cd","ef"], then "abcdef", "abefcd", "cdabef", "cdefab", "efabcd", and "efcdab"
// are all concatenated strings. "acdbef" is not a concatenated substring because it is not the concatenation of any permutation of words.
// Return the starting indices of all the concatenated substrings in s. You can return the answer in any order.

// Example 1:
// Input: s = "barfoothefoobarman", words = ["foo","bar"]
// Output: [0,9]
// Explanation: Since words.length == 2 and words[i].length == 3, the concatenated substring has to be of length 6.
// The substring starting at 0 is "barfoo". It is the concatenation of ["bar","foo"] which is a permutation of words.
// The substring starting at 9 is "foobar". It is the concatenation of ["foo","bar"] which is a permutation of words.
// The output order does not matter. Returning [9,0] is fine too.

// Example 2:
// Input: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
// Output: []
// Explanation: Since words.length == 4 and words[i].length == 4, the concatenated substring has to be of length 16.
// There is no substring of length 16 is s that is equal to the concatenation of any permutation of words.
// We return an empty array.

// Example 3:
// Input: s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
// Output: [6,9,12]
// Explanation: Since words.length == 3 and words[i].length == 3, the concatenated substring has to be of length 9.
// The substring starting at 6 is "foobarthe". It is the concatenation of ["foo","bar","the"] which is a permutation of words.
// The substring starting at 9 is "barthefoo". It is the concatenation of ["bar","the","foo"] which is a permutation of words.
// The substring starting at 12 is "thefoobar". It is the concatenation of ["the","foo","bar"] which is a permutation of words.

// Constraints:

// 1 <= s.length <= 104
// 1 <= words.length <= 5000
// 1 <= words[i].length <= 30
// s and words[i] consist of lowercase English letters.

// решение с использованием хэш-таблиц
// существует также более эффективное решение с использованием хэш-таблиц и скользящего окна
func findSubstring(s string, words []string) []int {
	var res = []int{}
	if len(words) == 0 || len(s) < len(words)*len(words[0]) {
		return res
	}

	size := len(words) * len(words[0])
	wm := make(map[string]int, len(words))
	for i := range words {
		wm[words[i]]++
	}

	for i := range s {
		if len(s)-i+1 <= size {
			break
		}
		if containsPermutations(s[i:i+size], wm, len(words[0])) {
			res = append(res, i)
		}
	}

	return res
}

func containsPermutations(s string, m map[string]int, length int) bool {
	log.Println(s)
	tm := make(map[string]int, len(m))

	for i := 0; i < len(s); i += length {
		tm[s[i:i+length]]++
	}
	log.Println(m, tm)
	for k, vi := range tm {
		if vj, ok := m[k]; !ok || vj != vi {
			return false
		}
	}

	return true
}
