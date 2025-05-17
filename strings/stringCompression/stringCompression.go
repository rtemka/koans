// #443
// https://leetcode.com/problems/string-compression/description/
package leetcode

import "strconv"

func compress(chars []byte) int {
	read, write := 0, 0
	n := len(chars)

	for read < n {
		char := chars[read]
		count := 0
		// Count all repeated letters.
		for read < n && chars[read] == char {
			read++
			count++
		}
		// Write the letter.
		chars[write] = char
		write++
		// Write the count.
		if count > 1 {
			d := strconv.Itoa(count)
			for j := range d {
				chars[write] = d[j]
				write++
			}
		}
	}
	return write
}
