package front_back

import "testing"

func TestMain(t *testing.T) {
	t.Log("Testing the output of front_back()")
	s := []string{"code", "a", "ab", "abc", "", "Chocolate", "aavJ", "hello"}
	r := []string{"eodc", "a", "ba", "cba", "", "ehocolatC", "Java", "oellh"}

	for i, v := range s {
		if front_back(v) != r[i] {
			t.Error("Expected output of %s, but it was %s instead.", r[i], front_back(v))
		}
	}
}
