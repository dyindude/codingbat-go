package string_times

import "testing"

func TestMain(t *testing.T) {
	t.Log("Testing output of string_times()")
	s := []string{"Hi", "Hi", "Hi", "Hi", "Hi", "Oh Boy!", "x", "", "code", "code"}
	n := []int{2, 3, 1, 0, 5, 2, 4, 4, 2, 3}
	r := []string{"HiHi", "HiHiHi", "Hi", "", "HiHiHiHiHi", "Oh Boy!Oh Boy!", "xxxx", "", "codecode", "codecodecode"}

	for i, v := range r {
		if string_times(s[i], n[i]) != v {
			t.Errorf("Expected output of %s, but got %s instead.", v, string_times(s[i], n[i]))
		}

	}
}
