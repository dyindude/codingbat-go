package front3

import "testing"

func TestMain(t *testing.T) {
	t.Log("Testing output of front3()")
	s := []string{"Java", "Chocolate", "abc", "abcXYZ", "ab", "a", ""}
	r := []string{"JavJavJav", "ChoChoCho", "abcabcabc", "abcabcabc", "ababab", "aaa", ""}
	for i, v := range r {
		if front3(s[i]) != v {
			t.Errorf("Expected output of %s, but it was %s instead.", v, s[i])
		}
	}
}
