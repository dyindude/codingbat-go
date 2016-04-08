package array_front9

import "testing"

func TestMain(t *testing.T) {
	t.Log("Testing output of array_front9")
	s := [][]int{
		{1, 2, 9, 3, 4},
		{1, 2, 3, 4, 9},
		{1, 2, 3, 4, 5},
		{9, 2, 3},
		{1, 9, 9},
		{1, 2, 3},
		{1, 9},
		{5, 5},
		{2},
		{9},
		//		{},
		{3, 9, 2, 3, 3},
	}

	r := []bool{
		true,
		false,
		false,
		true,
		true,
		false,
		true,
		false,
		false,
		true,
		//		false,
		true,
	}

	for i, v := range s {
		if array_front9(v) != r[i] {
			t.Errorf("Expected output of %s, but got %s instead.", r[i], array_front9(v))

		}
	}
}
