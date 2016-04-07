package string_splosion

func string_splosion(s string) string {
	var r string
	for i := 0; i < len(s)+1; i++ {
		r = r + s[:i]
	}
	return r
}
