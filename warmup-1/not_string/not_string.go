package not_string

func not_string(s string) string {
	if len(s) >= 3 && s[0:3] == "not" {
		return s
	} else {
		return "not " + s
	}
}
