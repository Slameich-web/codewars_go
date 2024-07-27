package kata

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	var result int = (dadYearsOld - sonYearsOld*2)
	if result < 0 {
		return -result
	}
	return result

}