package kata

func EvenOrOdd(number int) string {
	if number%2 == 1 || number%2 == -1 {
		return "Odd"
	} else {
		return "Even"
	}
}