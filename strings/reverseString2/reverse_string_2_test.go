package reversestring2

import "testing"

func Test_reverseStr(t *testing.T) {
	want := "bacdfeg"
	if got := reverseStr("abcdefg", 2); got != want {
		t.Fatalf("= %s, want %s", got, want)
	}

	want = "bacd"
	if got := reverseStr("abcd", 2); got != want {
		t.Fatalf("= %s, want %s", got, want)
	}
}
