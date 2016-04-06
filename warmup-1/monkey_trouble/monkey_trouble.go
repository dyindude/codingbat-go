package monkey_trouble

func monkey_trouble(a_smile bool, b_smile bool) bool {
	if (a_smile && b_smile) || (!a_smile && !b_smile) {
		return true
	} else {
		return false
	}
}
