// #787
// https://leetcode.com/problems/cheapest-flights-within-k-stops/description/
package leetcode

import (
	"math"
)

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	costs := make([]int, n)
	for i := range costs {
		costs[i] = math.MaxInt
	}
	costs[src] = 0

	flightCosts := make([]int, n)
	for range k + 1 {
		copy(flightCosts, costs)
		for _, flight := range flights {
			from, to, price := flight[0], flight[1], flight[2]
			if costs[from] != math.MaxInt {
				flightCosts[to] = min(flightCosts[to], costs[from]+price)
			}
		}
		copy(costs, flightCosts)
	}
	if costs[dst] == math.MaxInt {
		return -1
	} else {
		return costs[dst]
	}
}
