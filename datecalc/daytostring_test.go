package datecalc

import "testing"

func Test_Day_To_Second_Input_15103_Return_1304899200(t *testing.T) {
	expected := 1304899200
	actual := dayToSecond(15103)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
