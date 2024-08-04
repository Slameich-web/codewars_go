package kata

func GetCount(str string) (count int) {
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	for _, v := range str {
		for _, vowel := range vowels {
			if v == vowel {
				count++
			}
		}
	}
	return count
}