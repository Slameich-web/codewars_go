package kata

func Multiple3And5(number int) int {
	var result = 0
	for i := 0; i < number; i++ {
		if i%3 == 0 && i%5 == 0 {
			result += i
		} else if i%3 == 0 {
			result += i
		} else if i%5 == 0 {
			result += i
		}

	}
	return result
}