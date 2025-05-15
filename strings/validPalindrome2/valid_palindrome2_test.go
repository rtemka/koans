package validpalindrome2

import "testing"

func Test_validPalindrome(t *testing.T) {
	want := true
	if validPalindrome("aba") != want {
		t.Fatal("wrong", "aba")
	}

	want = true
	if validPalindrome("abca") != want {
		t.Fatal("wrong", "abca")
	}

	want = false
	if validPalindrome("abdwggfdca") != want {
		t.Fatal("wrong", "abdwggfdca")
	}

	want = false
	if validPalindrome("abc") != want {
		t.Fatal("wrong", "abc")
	}
}
