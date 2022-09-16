package leetcode

// You are given a string s and an integer k. You can choose any character of the string and change
// it to any other uppercase English character. You can perform this operation at most k times.

// Return the length of the longest substring containing the same letter you can get after performing the above operations.

// Example 1:
// Input: s = "ABAB", k = 2
// Output: 4
// Explanation: Replace the two 'A's with two 'B's or vice versa.

// Example 2:
// Input: s = "AABABBA", k = 1
// Output: 4
// Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
// The substring "BBBB" has the longest repeating letters, which is 4.

// Constraints:
// 1 <= s.length <= 105
// s consists of only uppercase English letters.
// 0 <= k <= s.length

// k == 2
// "ABBB"
//  ^
//     ^
// "BAAAB"

func characterReplacement(s string, k int) int {

	var max, mostFreq int   // максимальная длина и самый частый символ
	var lo, hi int          // указатели скользящего окна
	m := make(map[byte]int) // карта символов для определения самого частого

	for hi < len(s) {

		m[s[hi]]++
		this := m[s[hi]]
		// проверяем нужно ли обновить самый встречающийся символ
		if mostFreq < this {
			mostFreq = this
		}

		// сужаем окно, если понадобится заменить больше чем k символов
		if (hi-lo+1)-mostFreq > k {
			m[s[lo]] = m[s[lo]] - 1
			lo++
		}
		// проверяем новый ли у нас максимум длины
		if max < hi-lo+1 {
			max = hi - lo + 1
		}
		hi++
	}

	return max
}
