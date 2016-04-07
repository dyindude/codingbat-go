package makes10

import "testing"

func TestMain(t *testing.T) {
	t.Log("Testing output of makes10()")
	a := []int{9, 9, 1, 10, 10, 8, 8, 10, 12}
	b := []int{10, 9, 9, 1, 10, 2, 3, 42, -2}
	r := []bool{true, false, true, true, true, true, false, true, true}
	for i, _ := range r {
		if makes10(a[i], b[i]) != r[i] {
			t.Errorf("Expected result of %d, but got %d instead.", makes10(a[i], b[i]), r[i])
		}

	}
}
