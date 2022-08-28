package leetcode

import (
	"math"
)

// A perfect number is a positive integer that is equal to the sum of its positive divisors, excluding the number itself.
// A divisor of an integer x is an integer that can divide x evenly.

// Given an integer n, return true if n is a perfect number, otherwise return false.

// Example 1:

// Input: num = 28
// Output: true
// Explanation: 28 = 1 + 2 + 4 + 7 + 14
// 1, 2, 4, 7, and 14 are all divisors of 28.
// Example 2:

// Input: num = 7
// Output: false

func checkPerfectNumber(num int) bool {

	if num == 1 {
		return false
	}

	var sum int                            // сумма делителей
	var cap = int(math.Sqrt(float64(num))) // верхняя граница

	for i := 2; i <= cap; i++ {
		if sum > num {
			return false
		}

		// если делится нацело, то прибавляем к сумме делителей
		if num%i == 0 {
			sum += i + num/i
		}
	}

	return sum+1 == num
}
