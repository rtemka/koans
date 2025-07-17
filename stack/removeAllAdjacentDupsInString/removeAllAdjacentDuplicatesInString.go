// #1047
// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/description/
package leetcode

func removeDuplicates(s string) string {
	var stack []byte
	for i := range s {
		// If the current character is the same as the top character on the stack,
		// a pair of adjacent duplicates has been formed. So, pop the top character
		// from the stack.
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			// Otherwise, push the current character onto the stack.
			stack = append(stack, s[i])
		}
	}
	// Return the remaining characters as a string.
	return string(stack)
}
