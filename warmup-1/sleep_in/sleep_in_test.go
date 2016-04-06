package sleep_in

import "testing"

func TestMain(t *testing.T) {
	t.Log("Testing output of sleep_in()")
	arg1 := []bool{false, true, false, true}
	arg2 := []bool{false, false, true, true}
	r := []bool{true, false, true, true}
	for i, _ := range r {
		result := sleep_in(arg1[i], arg2[i])
		expect := r[i]
		if result != expect {
			t.Errorf("Expected output of %d, but it was %d instead.", expect, result)
		}
	}
}
