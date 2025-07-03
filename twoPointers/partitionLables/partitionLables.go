// #763
// https://leetcode.com/problems/partition-labels/description
package leetcode

func partitionLabels(s string) []int {
	var result []int
	// Store where each letter occurs for the last time.
	var lastIndex [26]int
	for i := range s {
		lastIndex[s[i]-'a'] = i
	}
	// Traverse the string again and keep track of the
	// farthest last occurrence you've seen so far.
	lo, hi := 0, 0
	for i := range s {
		hi = max(hi, lastIndex[s[i]-'a'])
		// If current index matches the 'hi' - cut the partition here.
		if i == hi {
			result = append(result, hi-lo+1)
			lo = i + 1
		}
	}

	return result
}

func partitionLabelsTwoSets(s string) []int {
	var result []int
	// Gather letters frequencies into the memo array.
	var total [26]byte
	for i := range s {
		total[s[i]-'a']++
	}
	// Define memo array to hold frequencies of partitioned letters.
	var partitioned [26]byte
	lo := 0
	for hi := range s {
		i := s[hi] - 'a'
		// Add to partitioned and subtract from total.
		partitioned[i]++
		total[i]--
		// If frequency of the letter in the total is decreased down to zero
		// it triggers the check if set of partitioned letters not intersects with
		// the set of total. If sets not intersects then add to result and move 'lo' pointer.
		if total[i] == 0 && !intersect(total, partitioned) {
			result = append(result, hi-lo+1)
			lo = hi + 1
		}
	}
	return result
}

func intersect(x, y [26]byte) bool {
	for i := range x {
		if x[i] > 0 && y[i] > 0 {
			return true
		}
	}
	return false
}
