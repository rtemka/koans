// #42
// https://leetcode.com/problems/valid-parentheses/description/
package leetcode

// Дана строка строки содержащая только символы '(', ')', '{', '}', '[' и ']', определите, является ли входная строка допустимой.
// Входная строка допустима, если:
// Открытые скобки должны быть закрыты скобками того же типа.
// Открытые скобки должны быть закрыты в правильном порядке.

// Example 1:

// Input: s = "()"
// Output: true
// Example 2:

// Input: s = "()[]{}"
// Output: true
// Example 3:

// Input: s = "(]"
// Output: false

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := make([]byte, 0, len(s)/2)

	for i := range s {
		switch {
		case s[i] == '(', s[i] == '{', s[i] == '[':
			stack = append(stack, s[i])
		case s[i] == ')' && len(stack) > 0 && stack[len(stack)-1] == '(':
			stack = stack[:len(stack)-1]
		case s[i] == '}' && len(stack) > 0 && stack[len(stack)-1] == '{':
			stack = stack[:len(stack)-1]
		case s[i] == ']' && len(stack) > 0 && stack[len(stack)-1] == '[':
			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}

	return len(stack) == 0
}

func IsValid2(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	var stack []byte
	m := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for i := range s {
		if _, ok := m[s[i]]; ok {
			stack = append(stack, s[i])
			continue
		}
		if len(stack) == 0 || m[stack[len(stack)-1]] != s[i] {
			return false
		}
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
