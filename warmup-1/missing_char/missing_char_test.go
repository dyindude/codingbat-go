package main

import "testing"

func TestMain(t *testing.T) {
	t.Log("Testing output of missing_char()")
	arg1 := []string{"kitten", "kitten", "kitten", "Hi", "Hi", "code", "code", "code", "code", "chocolate"}
	arg2 := []int{1, 0, 4, 0, 1, 0, 1, 2, 3, 8}
	r := []string{"ktten", "itten", "kittn", "i", "H", "ode", "cde", "coe", "cod", "chocolat"}
	for i, _ := range r {
		result := missing_char(arg1[i], arg2[i])
		expect := r[i]
		if result != expect {
			t.Errorf("Expected output of %d, but it was %d instead.", expect, result)
		}
	}
}
