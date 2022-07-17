package leetcode

import "strings"

// Given an array of strings words, return the words that can be typed
// using letters of the alphabet on only one row of American keyboard like the image below.
// In the American keyboard:
// the first row consists of the characters "qwertyuiop",
// the second row consists of the characters "asdfghjkl", and
// the third row consists of the characters "zxcvbnm".

// Example 1:

// Input: words = ["Hello","Alaska","Dad","Peace"]
// Output: ["Alaska","Dad"]
// Example 2:

// Input: words = ["omk"]
// Output: []
// Example 3:

// Input: words = ["adsdf","sfd"]
// Output: ["adsdf","sfd"]

func findWords(words []string) []string {

	var vwords []string
	for i := range words {

		var c int
		var valid bool

		for _, r := range strings.ToLower(words[i]) {
			switch r {
			case 'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p':
				if c == 0 {
					valid = true
					c = 1
				} else if c != 1 {
					valid = false
					break
				}
			case 'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l':
				if c == 0 {
					valid = true
					c = 2
				} else if c != 2 {
					valid = false
					break
				}
			case 'z', 'x', 'c', 'v', 'b', 'n', 'm':
				if c == 0 {
					valid = true
					c = 3
				} else if c != 3 {
					valid = false
					break
				}
			default:
				valid = false
				break
			}
		}
		if valid {
			vwords = append(vwords, words[i])
		}
	}

	return vwords
}
