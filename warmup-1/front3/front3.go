package front3

func front3(s string) string {
	var front string
	if len(s) < 3 {
		front = s[:len(s)]
	} else {
		front = s[:3]
	}
	return front + front + front
}
