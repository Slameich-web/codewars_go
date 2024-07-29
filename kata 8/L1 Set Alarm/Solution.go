package kata

func SetAlarm(employed, vacation bool) bool {
	if employed && vacation == false {
		return true
	}
	return false
}