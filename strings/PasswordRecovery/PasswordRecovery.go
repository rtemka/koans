package main

import (
	"crypto/md5"
	"fmt"
)

var alphabet = []rune{'a', 'b', 'c', 'd', '1', '2', '3'}

func RecoverPassword(h []byte) string {
	var passwords []string
	for _, alph := range alphabet {
		passwords = append(passwords, string(alph))
	}
	// fmt.Println(passwords)

	for i := 0; i < 1; i++ {
		for _, p := range passwords {
			hh := hashPassword(p)
			if string(hh) == string(h) {
				return string(hh)
			}

		}
		passwords = addPass(passwords)
	}
	return ""
}

func addPass(ss []string) []string {
	out := make([]string, 0, len(ss)*len(alphabet))
	for _, s := range ss {
		for _, r := range alphabet {
			out = append(out, s+fmt.Sprint(r))
		}
	}
	return out
}

func hashPassword(in string) []byte {
	h := md5.Sum([]byte(in))
	return h[:]
}
