// #424
// https://leetcode.com/problems/longest-repeating-character-replacement/description
package leetcode

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

func characterReplacement2025(s string, k int) int {
	// Intuition:
	// The key observation is that the minimum number of replacements needed to achieve uniformity is
	// obtained by replacing all characters except the most frequent one.
	//
	// If num_chars_to_replace <= k, the substring can be made unirform
	// and we can EXPAND sliding window(that is advance the right pointer).
	//
	// If num_chars_to_replace > k, the substring cannot be made uniform
	// and we MUST SLIDE sliding window.(that is advance left AND right pointers).

	freqs := make(map[byte]int)
	maxLen, highestFreq := 0, 0
	left, right := 0, 0
	n := len(s)
	for right < n {
		// Update the frequency of the character at the right pointer
		// and the highest frequency for the current window.
		freqs[s[right]]++
		highestFreq = max(highestFreq, freqs[s[right]])
		// Calculate replacements needed for the current Window.
		numCharsToReplace := (right - left + 1) - highestFreq
		// Slide the window if the number of replacements needed exceeds
		// 'k '. The right pointer always gets advanced, so we just need
		// to advance 'left'.
		if numCharsToReplace > k {
			// Remove the character at the left pointer from the hash map
			// before advancing the left pointer.
			freqs[s[left]]--
			left++
		}
		// Since the length of the current window increases or stays the
		// same, assign the length of the current window to 'maxLen'.
		maxLen = right - left + 1
		// Expand the window.
		right++
	}
	return maxLen
}
