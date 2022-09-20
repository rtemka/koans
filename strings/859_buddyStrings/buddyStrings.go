package leetcode

// Given two strings s and goal, return true if you can swap two letters in s so the result is equal to goal, otherwise, return false.

// Swapping letters is defined as taking two indices i and j (0-indexed) such that i != j and swapping the characters at s[i] and s[j].

// For example, swapping at indices 0 and 2 in "abcd" results in "cbad".

// Example 1:

// Input: s = "ab", goal = "ba"
// Output: true
// Explanation: You can swap s[0] = 'a' and s[1] = 'b' to get "ba", which is equal to goal.
// Example 2:

// Input: s = "ab", goal = "ab"
// Output: false
// Explanation: The only letters you can swap are s[0] = 'a' and s[1] = 'b', which results in "ba" != goal.
// Example 3:

// Input: s = "aa", goal = "aa"
// Output: true
// Explanation: You can swap s[0] = 'a' and s[1] = 'a' to get "aa", which is equal to goal.

// Constraints:
// 1 <= s.length, goal.length <= 2 * 104
// s and goal consist of lowercase letters.

// "abcd"
// "cbad"
// "dacb"

func buddyStrings(s string, goal string) bool {

	// если не одинаковая длина, то точно нет
	if len(s) != len(goal) {
		return false
	}

	// если строки одинаковые, и есть хоть одна буква
	// которая повторяется более одного раза, то можно
	// свапнуть, иначе нет
	if s == goal {
		seen := make(map[byte]bool)
		for i := range s {
			if seen[s[i]] {
				return true
			}
			seen[s[i]] = true
		}
		return false
	}

	// находим индексы разницы, в которых строки
	// отличаются. Их должно быть ровно два. Иначе false
	var first, second = -1, -1
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			if first == -1 {
				first = i
			} else if second == -1 {
				second = i
			} else {
				return false
			}
		}
	}

	// проверяем был ли найден второй индекс, и если да, то равны ли символы
	// который предназначены для свапа
	return second != -1 && s[first] == goal[second] && s[second] == goal[first]
}
