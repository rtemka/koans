// #9
// https://leetcode.com/problems/palindrome-number/description/
package leetcode

func isPalindrome(x int) bool {
	y, n := 0, x
	for n > 0 {
		y = y*10 + n%10
		n /= 10
	}
	return y == x
}
