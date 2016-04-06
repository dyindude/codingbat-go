package parrot_trouble

func parrot_trouble(talking bool, hour int) bool {
	if talking && (hour < 7 || hour > 20) {
		return true
	} else {
		return false
	}
}
