package leetcode

import (
	"fmt"
	"math"
	"strconv"
)

// Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed
// 32-bit integer range [-2^31, 2^31 - 1], then return 0.

// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

// Example 1:

// Input: x = 123
// Output: 321
// Example 2:

// Input: x = -123
// Output: -321
// Example 3:

// Input: x = 120
// Output: 21

func reverse(x int) int {
	var res int

	for i := x; i != 0; i /= 10 {
		if i != x {
			res = res*10 + i%10
		} else {
			res = i % 10
		}

		if res > math.MaxInt32 || res < math.MinInt32 {
			return 0
		}
	}

	return res
}

func reverseStr(x int) int {

	s := []rune(fmt.Sprint(x))

	i := 0
	if x < 0 {
		i = 1
	}
	for j := len(s) - 1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	res, err := strconv.Atoi(string(s))

	if err == nil && res <= math.MaxInt32 && res >= math.MinInt32 {
		return res
	} else {
		return 0
	}
}
