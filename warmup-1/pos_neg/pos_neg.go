package pos_neg

func pos_neg(a int, b int, negative bool) bool {
	if negative {
		return (a < 0 && b < 0)
	} else {
		return ((a < 0 && b > 0) || (a > 0 && b < 0))
	}
}
