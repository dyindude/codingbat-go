package sum_double

import "testing"

func TestMain(t *testing.T) {
	t.Log("Testing output of sum_double()")
	a := []int{1, 3, 2, -1, 3, 0, 0, 3}
	b := []int{2, 2, 2, 0, 3, 0, 1, 4}
	r := []int{3, 5, 8, -1, 12, 0, 1, 7}

	for i, _ := range r {
		result := sum_double(a[i], b[i])
		expect := r[i]
		if result != expect {
			t.Errorf("Expected output of %d, but it was %d instead.", expect, result)
		}

	}

}
