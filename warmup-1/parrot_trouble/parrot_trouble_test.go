package parrot_trouble

import "testing"

func TestMain(t *testing.T) {
	t.Log("Testing output of parrot_trouble()")
	arg1 := []bool{true, true, false, true, false, false, true, false, true, false}
	arg2 := []int{6, 7, 6, 21, 21, 20, 23, 23, 20, 12}
	r := []bool{true, false, false, true, false, false, true, false, false, false}
	for i, _ := range r {
		result := parrot_trouble(arg1[i], arg2[i])
		expect := r[i]
		if result != expect {
			t.Errorf("Expected output of %d, but it was %d instead.", expect, result)
		}
	}
}
