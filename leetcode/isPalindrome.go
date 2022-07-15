package leetcode

import "fmt"

func isPalindrome(x int) bool {
	xs := fmt.Sprint(x)
	for i, j := 0, len(xs)-1; i < len(xs)/2; i, j = i+1, j-1 {
		if xs[i] != xs[j] {
			return false
		}
	}
	return true
}
