// #66
// https://leetcode.com/problems/plus-one/description/
package plusone

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+carry == 10 {
			digits[i] = 0
		} else {
			digits[i] += carry
			carry = 0
			break
		}
	}
	if carry > 0 {
		return append([]int{1}, digits...)
	}
	return digits
}
