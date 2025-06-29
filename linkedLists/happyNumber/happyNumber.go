// #202
// https://leetcode.com/problems/happy-number/description/
package leetcode

func isHappy(n int) bool {
	// If the fast and slow pointers meet, it indicates the presence of a cycle, meaning
	// the number is not a happy number. Otherwise, the fast pointer will reach 1, in which
	// case the number is a happy number.
	slow, fast := n, n
	for {
		// Advance slow pointer to the next number in sequence.
		slow = getNext(slow)
		// Advance fast pointer 2 numbers ahead in sequence.
		fast = getNext(getNext(fast))
		if fast == 1 {
			return true
		} else if slow == fast {
			return false
		}
	}
}

func getNext(n int) int {
	next := 0
	for n > 0 {
		// Get the last digit out of 'n'.
		lastDigit := n % 10
		// Add square of the last digit to sum.
		next += lastDigit * lastDigit
		// Remove last digit from 'n'.
		n /= 10
	}
	return next
}
