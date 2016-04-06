package near_hundred

import "math"

func near_hundred(n int) bool {
	if (math.Abs(100-float64(n)) <= 10) || (math.Abs(200-float64(n)) <= 10) {
		return true
	} else {
		return false
	}
}
