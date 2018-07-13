package datecalc

func calculateCommonYears(days int) float64 {
	daysToFloat := float64(days)
	result := float64((daysToFloat / 365.0) * 100.0)
	return result
}
