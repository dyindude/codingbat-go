package monkey_trouble

import "testing"

func TestMain(t *testing.T) {
	t.Log("Testing output of monkey_trouble()")
	arg1 := []bool{true, false, true, false}
	arg2 := []bool{true, false, false, true}
	r := []bool{true, true, false, false}
	for i, _ := range r {
		result := monkey_trouble(arg1[i], arg2[i])
		expect := r[i]
		if result != expect {
			t.Errorf("Expected output of %d, but it was %d instead.", expect, result)
		}
	}
}
