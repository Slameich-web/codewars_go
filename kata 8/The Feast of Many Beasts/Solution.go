package kata

func Feast(beast string, dish string) bool {
	first := beast[0:1]
	last := beast[len(beast)-1:]
	if first == dish[0:1] && last == dish[len(dish)-1:] {
		return true
	}
	return false

}