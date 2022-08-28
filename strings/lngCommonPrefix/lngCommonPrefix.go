package leetcode

// Напишите функцию для поиска самой длинной строки общего префикса среди массива строк.
// Если общего префикса нет, верните пустую строку "".

// Example 1:

// Input: strs = ["flower","flow","flight"]
// Output: "fl"
// Example 2:

// Input: strs = ["dog","race car","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return ""
	}
	var l int
	xs := strs[1:]

	for i := range strs[0] {
		for j := range xs {
			if i >= len(xs[j]) || strs[0][i] != xs[j][i] {
				return strs[0][:l]
			}
		}
		l++
	}

	return strs[0][:l]
}
