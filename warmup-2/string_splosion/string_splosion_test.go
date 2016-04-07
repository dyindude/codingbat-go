package string_splosion

import "testing"

func TestMain(t *testing.T) {
	t.Log("Testing output of string_splosion()")
	s := []string{
		"Code",
		"abc",
		"ab",
		"x",
		"fade",
		"There",
		"Kitten",
		"Bye",
		"Good",
		"Bad",
	}
	r := []string{
		"CCoCodCode",
		"aababc",
		"aab",
		"x",
		"ffafadfade",
		"TThTheTherThere",
		"KKiKitKittKitteKitten",
		"BByBye",
		"GGoGooGood",
		"BBaBad",
	}
	for i, v := range r {
		if string_splosion(s[i]) != v {
			t.Errorf("Expected output of %s, but got %s instead.", v, string_splosion(s[i]))
		}
	}
}
