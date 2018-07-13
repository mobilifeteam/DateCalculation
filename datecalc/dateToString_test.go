package datecalc

import "testing"

func TestDateToStringInputDateShouldReturnString(t *testing.T) {
	expected := "7/3/1977"
	actual := DateToString(7, 3, 1977)

	if expected != actual {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}
