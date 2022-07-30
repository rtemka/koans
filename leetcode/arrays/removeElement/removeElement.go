package leetcode

// Given an integer array nums and an integer val, remove
// all occurrences of val in nums in-place. The relative order of the elements may be changed.

// Since it is impossible to change the length of the array in some languages,
// you must instead have the result be placed in the first part of the array nums.
// More formally, if there are k elements after removing the duplicates,
// then the first k elements of nums should hold the final result.
// It does not matter what you leave beyond the first k elements.

// Return k after placing the final result in the first k slots of nums.

// Do not allocate extra space for another array. You must do this
// by modifying the input array in-place with O(1) extra memory.

func removeElementNaive(nums []int, val int) int {
	k := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			k--
			i--
		}
	}
	return k
}

func removeElement(nums []int, val int) int {
	var j int = 0

	for i, n := range nums {
		if n != val {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}
