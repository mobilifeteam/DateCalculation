package datecalc

import "testing"

func Test_DaystoHours_Input_15103_Should_Be_362472(t *testing.T){
	expected := 362472
	actual := daysToHours(15103)

	if expected != actual{
		t.Errorf("expected %d but got %d",expected,actual)
	}
}