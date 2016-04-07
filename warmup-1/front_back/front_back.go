package front_back

func front_back(s string) string {
	if len(s) <= 1 {
		return s
	} else {
		return s[len(s)-1:] + s[1:len(s)-1] + s[0:1]
	}
}
