package leetcode

func TwoSum(nums []int, target int) []int {
	for i, j := 1, 0; i < len(nums); i, j = i+1, j+1 {
		if nums[i]+nums[j] == target {
			return []int{j, i}
		}
	}
	return []int{-1, -1}
}
