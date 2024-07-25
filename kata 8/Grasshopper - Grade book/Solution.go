package kata

func GetGrade(a, b, c int) rune {
	result := (a + b + c) / 3
	if result >= 90 {
		return 'A'
	}
	if result >= 80 {
		return 'B'
	}
	if result >= 70 {
		return 'C'
	}
	if result >= 60 {
		return 'D'
	}
	return 'F'
}