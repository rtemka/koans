package leetcode

// Целочисленный массив называется арифметическим, если он состоит как минимум из трех элементов
// и если разница между любыми двумя последовательными элементами одинакова.

// Например, [1,3,5,7,9], [7,7,7,7] и [3,-1,-5,-9] являются арифметическими последовательностями.
// Учитывая целочисленный массив nums, вернуть количество арифметических подмассивов nums.

// Подмассив — это непрерывная подпоследовательность массива.

// Пример 1:

// Ввод: nums = [1,2,3,4]
// Выход: 3
// Объяснение: у нас есть 3 арифметических среза в nums: [1, 2, 3], [2, 3, 4] и собственно [1,2,3,4].

// Пример 2:

// Ввод: nums = [1]
// Выход: 0

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	var count, ind int
	var pdiff = nums[1] - nums[0]

	for i := 1; i < len(nums)-1; i++ {
		diff := nums[i+1] - nums[i]
		if diff == pdiff {
			ind++
		} else {
			pdiff = diff
			ind = 0
		}
		count += ind
	}

	return count
}
