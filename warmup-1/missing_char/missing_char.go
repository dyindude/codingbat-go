package missing_char

func missing_char(s string, n int) string {
	return s[:n] + s[n+1:]
}
