package not_string

import "testing"

func TestMain(t *testing.T) {
	t.Log("Testing output of not_string()")
	s := []string{"candy", "x", "not bad", "bad", "not", "is not", "no"}
	r := []string{"not candy", "not x", "not bad", "not bad", "not", "not is not", "not no"}
	for i, v := range r {
		if not_string(s[i]) != v {
			t.Error("Expected output of %s, but it was %s instead.", v, s[i])
		}
	}
}
