package kata

import "strconv"

func BonusTime(salary int, bonus bool) string {
	if bonus {
		return "£" + strconv.Itoa(salary * 10)
	}
	return "£" + strconv.Itoa(salary)
}