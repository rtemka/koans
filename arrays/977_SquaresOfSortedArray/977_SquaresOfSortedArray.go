package leetcode

import (
	"log"
	"sort"
)

// Given an integer array nums sorted in non-decreasing order, return an array of
// the squares of each number sorted in non-decreasing order.

// Example 1:
// Input: nums = [-4,-1,0,3,10]
// Output: [0,1,9,16,100]
// Explanation: After squaring, the array becomes [16,1,0,9,100].
// After sorting, it becomes [0,1,9,16,100].

// Example 2:
// Input: nums = [-7,-3,2,3,11]
// Output: [4,9,9,49,121]

// Constraints:
// 1 <= nums.length <= 104
// -104 <= nums[i] <= 104
// nums is sorted in non-decreasing order.

// Follow up: Squaring each element and sorting the new array is very trivial,
// could you find an O(n) solution using a different approach?

// [-2,-3,-6,-7,0,1,4,5,10]
// [0,1,4,-7,-2,-3,-6,5,10]
// 	       ^
//                    ^
// [-4,-1,0,3,10]
// [3,-1,0,16,100]
// 	     ^
//  ^
//
// [-7,-6,-5,-4,-3,0,2,8]
// [2,0,9,16,25,36,49,64]
// [0,2,-3,-4,-5,-6,-7]
//	^
//                     ^
//
// [-7,-6,-5,-4,-2,-3,-1]
// [1,-6,-5,-4,-2,-3,-7]
//	^
//                 ^

// легкое решение
func sortedSquares(nums []int) []int {
	for i := range nums {
		nums[i] = nums[i] * nums[i]
	}
	sort.Ints(nums)
	return nums
}

// методом двух указателей с нахождением минимального значения
func sortedSquares2(nums []int) []int {

	var j int
	for j < len(nums) && nums[j] < 0 {
		j++
	}

	var res = make([]int, len(nums))

	for k, i := 0, j-1; j-i <= len(nums); k++ {
		switch {
		case j == len(nums):
			res[k] = nums[i] * nums[i]
			i--
		case i < 0:
			res[k] = nums[j] * nums[j]
			j++
		case -nums[i] < nums[j]:
			res[k] = nums[i] * nums[i]
			i--
		default:
			res[k] = nums[j] * nums[j]
			j++
		}
	}
	return res
}

// методом двух указателей с разных концов массива
func sortedSquares2Ends(nums []int) []int {
	var lo, hi = 0, len(nums) - 1
	var k = hi

	var res = make([]int, len(nums))

	for k >= 0 {
		left := nums[lo] * nums[lo]
		right := nums[hi] * nums[hi]
		if left > right {
			lo++
			res[k] = left
		} else {
			hi--
			res[k] = right
		}
		k--
	}

	return res
}

func reverseSquare(nums []int) {
	// log.Print(nums)
	mid := len(nums) / 2
	for i := 0; i < mid; i++ {
		j := len(nums) - 1 - i
		nums[i], nums[j] = nums[j]*nums[j], nums[i]*nums[i]
	}
	if mid > 0 && len(nums)%2 != 0 {
		nums[mid] = nums[mid] * nums[mid]
	}
	log.Print(nums)
}

// // с использованием двух указателей
// func sortedSquares2(nums []int) []int {
// 	// [0,-1,-4,3,10]
// 	// 	   ^
// 	//          ^
// 	for i, j := 0, 1; i < len(nums); {
// 		log.Print(nums, i, j)
// 		swaped := false
// 		if i >= j {
// 			nums[i] = nums[i] * nums[i]
// 			i++
// 			continue
// 		}

// 		if abs(nums[i]) > abs(nums[j]) {
// 			nums[i], nums[j] = nums[j], nums[i]
// 			swaped = true
// 		}
// 		if nums[i] >= 0 || (!swaped && nums[j] >= 0) || !swaped {
// 			nums[i] = nums[i] * nums[i]
// 			i++
// 		} else if j+1 == len(nums) {
// 			j--
// 		} else {
// 			j++
// 		}

// 		log.Println(swaped)
// 	}

// 	return nums
// }

// func abs(x int) int {
// 	if x < 0 {
// 		return -x
// 	}
// 	return x
// }
