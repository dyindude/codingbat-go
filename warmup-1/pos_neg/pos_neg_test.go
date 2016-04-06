package pos_neg

import "testing"

func TestMain(t *testing.T) {
	t.Log("Testing output of pos_neg()")
	arg1 := []int{1, -1, -4, -4, -4, -4, 1, -1, 1, -1, 1, -1, 5, -6, -5, -2, 1, -5, -5}
	arg2 := []int{-1, 1, -5, -5, 5, 5, 1, -1, -1, 1, 1, -1, -5, 6, -6, -1, 2, 6, -5}
	arg3 := []bool{false, false, true, false, false, true, false, false, true, true, true, true, false, false, false, false, false, true, true}
	r := []bool{true, true, true, false, true, false, false, false, false, false, false, true, true, true, false, false, false, false, true}
	for i, _ := range r {
		result := pos_neg(arg1[i], arg2[i], arg3[i])
		expect := r[i]
		if result != expect {
			t.Error("Expected output of %d, but it was %d instead.", expect, result)
		}
	}

}
