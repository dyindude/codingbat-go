package near_hundred

import "testing"

func TestMain(t *testing.T) {
	t.Log("Testing output of near_hundred()")
	arg1 := []int{93, 90, 89, 110, 111, 121, 0, 5, 191, 189, 190, 200, 210, 211, 290}
	r := []bool{true, true, false, true, false, false, false, false, true, false, true, true, true, false, false}
	for i, _ := range r {
		result := near_hundred(arg1[i])
		expect := r[i]
		if result != expect {
			t.Errorf("Expected output of %d, but it was %d instead.", expect, result)
		}
	}
}
