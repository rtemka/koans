// #13
// https://leetcode.com/problems/roman-to-integer/description/
package leetcode

func romanToInt(s string) int {
	romans := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	num := 0
	for i := 0; i < len(s); i++ {
		num += romans[s[i]]
		if i > 0 {
			if romans[s[i-1]] < romans[s[i]] {
				num -= 2 * romans[s[i-1]]
			}
		}
	}
	return num
}

func RomanToInt2022(s string) int {

	var r int

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'I':
			r++
			if i+1 < len(s) && s[i+1] == 'V' {
				r += 3
				i++
			}
			if i+1 < len(s) && s[i+1] == 'X' {
				r += 8
				i++
			}
		case 'V':
			r += 5
		case 'X':
			r += 10
			if i+1 < len(s) && s[i+1] == 'L' {
				r += 30
				i++
			}
			if i+1 < len(s) && s[i+1] == 'C' {
				r += 80
				i++
			}
		case 'L':
			r += 50
		case 'C':
			r += 100
			if i+1 < len(s) && s[i+1] == 'D' {
				r += 300
				i++
			}
			if i+1 < len(s) && s[i+1] == 'M' {
				r += 800
				i++
			}
		case 'D':
			r += 500
		case 'M':
			r += 1000
		}
	}

	return r
}
