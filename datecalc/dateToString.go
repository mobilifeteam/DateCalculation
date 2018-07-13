package datecalc

import "strconv"

func dateToString(days int, months int, years int) string {
	return strconv.Itoa(days) + "/" + strconv.Itoa(months) + "/" + strconv.Itoa(years)
}
