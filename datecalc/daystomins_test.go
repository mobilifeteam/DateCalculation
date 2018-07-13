package datecalc

import "testing"

func Test_DaystoMins_Input_15103_Should_Be_21748320(t *testing.T){
	expected := 21748320
	actual := daysToMins(15103)

	if expected != actual{
		t.Errorf("expected %d but got %d",expected,actual)
	}
}