package leetcode

// Учитывая массив целых чисел nums и целочисленное целевое значение, верните
// индексы двух чисел таким образом, чтобы они складывались в целевое значение.
// Вы можете предположить, что каждый ввод будет иметь ровно одно решение, и
// вы не можете использовать один и тот же элемент дважды.
// Вы можете вернуть ответ в любом порядке.

// Example 1:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:

// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:

// Input: nums = [3,3], target = 6
// Output: [0,1]

// без дополнительной памяти
func TwoSum(nums []int, target int) []int {
	for i, j := 0, 1; j < len(nums); {
		if nums[i]+nums[j] == target {
			return []int{i, j}
		} else if j+1 == len(nums) {
			i++
			j = i + 1
		} else {
			j++
		}
	}
	return []int{-1, -1}
}

// с прмененние карты
func twoSum(nums []int, target int) []int {
	hash := map[int]int{}

	for i, num := range nums {
		if value, ok := hash[target-num]; ok {
			return []int{value, i}
		}
		hash[num] = i
	}
	return []int{}
}
