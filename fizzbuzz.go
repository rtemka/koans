package leetcode

import "fmt"

func FizzBuzz(n int) []string {
	answer := make([]string, 0, n)

	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			answer = append(answer, "FizzBuzz")
		case i%5 == 0:
			answer = append(answer, "Buzz")
		case i%3 == 0:
			answer = append(answer, "Fizz")
		default:
			answer = append(answer, fmt.Sprint(i))
		}
	}
	return answer
}
