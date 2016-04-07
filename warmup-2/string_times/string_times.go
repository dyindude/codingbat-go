package string_times

func string_times(s string, n int) string {
	var r string
	for i := 1; i <= n; i++ {
		r += s
	}
	return r
}
