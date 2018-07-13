package datecalc

import "testing"

func Test_Date_To_String_Input_Date_Int_7_3_1977_Return_String_731977(t *testing.T) {
	expected := "7/3/1977"
	actual := dateToString(7, 3, 1977)

	if expected != actual {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}
