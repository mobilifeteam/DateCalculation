package datecalc

import "strconv"

func DateToString(day int, month int, year int) string {
	return strconv.Itoa(day) + "/" + strconv.Itoa(month) + "/" + strconv.Itoa(year)
}
