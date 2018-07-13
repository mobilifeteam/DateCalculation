package datecalc

func dayToWeek(days int) (int, int) {
	return days / 7, days % 7
}
