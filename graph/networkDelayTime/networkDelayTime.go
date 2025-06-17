// #743
// https://leetcode.com/problems/network-delay-time/description/
package leetcode

import (
	"math"
)

// Bellman-Ford
func networkDelayTime(times [][]int, n int, k int) int {
	distanceTo := make([]int, n)
	for i := range distanceTo {
		distanceTo[i] = math.MaxInt
	}
	distanceTo[k-1] = 0

	for range n - 1 {
		for i := range times {
			// Relaxation.
			from, to, delay := times[i][0]-1, times[i][1]-1, times[i][2]
			if distanceTo[from] != math.MaxInt {
				distanceTo[to] = min(distanceTo[to], distanceTo[from]+delay)
			}
		}
	}
	// Granted signal distance from source is 0.
	longestSignalDistance := 0
	for i := range distanceTo {
		// If any vertex is unreachable return -1.
		if distanceTo[i] == math.MaxInt {
			return -1
		}
		longestSignalDistance = max(longestSignalDistance, distanceTo[i])
	}

	return longestSignalDistance
}
