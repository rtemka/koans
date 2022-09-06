package leetcode

// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

// Note that you must do this in-place without making a copy of the array.

// Example 1:
// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]

// Example 2:
// Input: nums = [0]
// Output: [0]

// Constraints:

// 1 <= nums.length <= 104
// -231 <= nums[i] <= 231 - 1

// Follow up: Could you minimize the total number of operations done?

func MoveZeroes(nums []int) {
	// объявляем два указателя, первый ищет нули, второй не нули
	for i, j := 0, 1; j < len(nums); {

		if nums[i] != 0 {
			i, j = i+1, j+1 // двигаем оба
		} else if nums[j] == 0 {
			j++ // двигаем только правый
		} else {
			nums[i], nums[j] = nums[j], nums[i] // свапаем
		}
	}
}

// func MoveZeroes2(nums []int) {
// 	// объявляем два указателя, первый ищет нули, второй не нули
// 	for i, j := 0, 1; j < len(nums); {
// 		if j == i || nums[j] == 0 {
// 			j++
// 			continue
// 		}

// 		if nums[i] == 0 {
// 			nums[i], nums[j] = nums[j], nums[i]
// 			i, j = i+1, j+1
// 		} else {
// 			i++
// 		}
// 	}
// }
