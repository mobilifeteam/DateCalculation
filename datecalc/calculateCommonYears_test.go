package datecalc

import "testing"

func TestCalculateCommonYearsInput15103ShouldReturn4137dot81Percen(t *testing.T) {
	expected := float64(4137.81)
	actual := calculateCommonYears(15103)

	if (expected - actual) < 0.001 {
		t.Errorf("expected %.2f but got %.2f", expected, actual)
	}
}
