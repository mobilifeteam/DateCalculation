package datecalc

import "testing"

func Test_dayToWeek(t *testing.T) {
	expectedWeeks := 2157
	expectedDays := 4
	actualWeeks, actualDays := dayToWeek(15103)

	if expectedWeeks != actualWeeks && expectedDays != actualDays {
		t.Errorf("Expected %d, %d but got %d %d", expectedWeeks, expectedDays, actualWeeks, actualDays)
	}
}
