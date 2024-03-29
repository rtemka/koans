package leetcode

import (
	"math"
)

// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints
// of the ith line are (i, 0) and (i, height[i]).

// Find two lines that together with the x-axis form a container, such that the container contains the most water.

// Return the maximum amount of water a container can store.

// Notice that you may not slant the container.

// Example 1:

// Input: height = [1,8,6,2,5,4,8,3,7]
// Output: 49
// Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7].
// In this case, the max area of water (blue section) the container can contain is 49.
// Example 2:

// Input: height = [1,1]
// Output: 1

// Constraints:

// n == height.length
// 2 <= n <= 10^5
// 0 <= height[i] <= 10^4
// Accepted
// 1,702,330
// Submissions
// 3,144,879

func maxAreaSlow(height []int) int {

	var max, area, dist int

	for i := range height {
		for j := range height {
			if j > i {
				dist = j - i
			} else {
				dist = i - j
			}

			if height[j] > height[i] {
				area = height[i] * dist
			} else {
				area = height[j] * dist
			}

			if area > max {
				max = area
			}
		}
	}

	return max
}

func maxArea(height []int) int {

	var length = len(height)
	var lo, hi = 0, length - 1
	var maxArea = math.MinInt

	for lo < hi {
		area := (hi - lo) * min(height[lo], height[hi])
		if area > maxArea {
			maxArea = area
		}
		if height[lo] < height[hi] {
			lo++
		} else {
			hi--
		}
	}

	return maxArea
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
