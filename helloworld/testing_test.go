package helloworld

import "testing"

func TestStringReturn(t *testing.T) {
	expected := "Hello World"
	res := StringReturn()
	if res != expected {
		t.Errorf("StringReturn returned an incorrect value expected %s but got %s", expected, res)
	}
}
