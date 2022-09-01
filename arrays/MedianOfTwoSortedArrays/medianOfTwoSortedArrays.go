package leetcode

import "log"

// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

// The overall run time complexity should be O(log (m+n)).

// Example 1:

// Input: nums1 = [1,3], nums2 = [2]
// Output: 2.00000
// Explanation: merged array = [1,2,3] and median is 2.
// Example 2:

// Input: nums1 = [1,2], nums2 = [3,4]
// Output: 2.50000
// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

// Constraints:

// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -10^6 <= nums1[i], nums2[i] <= 10^6

// [1,3,5,7]
// [2,4,6,8]
// [1,2,3,4,5,6,7,8]

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	switch {
	case len(nums1)+len(nums2) == 0:
		return 0
	case len(nums2) == 0:
		return median(nums1)
	case len(nums1) == 0:
		return median(nums2)
	case nums1[len(nums1)-1] < nums2[0]:
		nums1 = append(nums1, nums2...)
		return median(nums1)
	case nums2[len(nums2)-1] < nums1[0]:
		nums2 = append(nums2, nums1...)
		return median(nums2)
	}

	lo := SearchInsert(nums1, nums2[0])
	merged := make([]int, 0, len(nums1)+len(nums2))

	merged = append(merged, nums1[:lo]...)

	for i, j := lo, 0; ; {
		if j >= len(nums2) {
			merged = append(merged, nums1[i:]...)
			break
		}
		if i >= len(nums1) {
			merged = append(merged, nums2[j:]...)
			break
		}
		if nums2[j] < nums1[i] {
			merged = append(merged, nums2[j])
			j++
		} else {
			merged = append(merged, nums1[i])
			i++
		}
	}

	return median(merged)
}

func SearchInsert(nums []int, target int) int {

	lo := 0
	hi := len(nums) - 1

	for lo < hi {
		mid := (lo + hi) / 2

		diff := target - nums[mid]
		if diff < 0 {
			hi = mid - 1
		} else if diff > 0 {
			lo = mid + 1
		} else {
			return mid
		}
	}

	if target > nums[lo] {
		return lo + 1
	} else {
		return lo
	}
}

func median(arr []int) float64 {
	if len(arr)%2 != 0 {
		log.Println(arr[len(arr)/2])
		return float64(arr[len(arr)/2])
	}
	return float64(arr[len(arr)/2]+arr[len(arr)/2-1]) / 2
}
