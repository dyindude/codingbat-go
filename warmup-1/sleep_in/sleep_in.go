package sleep_in

func sleep_in(weekday bool, vacation bool) bool {
	if weekday != true || vacation {
		return true
	} else {
		return false
	}
}
