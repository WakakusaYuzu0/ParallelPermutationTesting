package example

import "testing"

func TestExample(t *testing.T) {
	ppt := Parapermtest{
		isDryRun:      false,
		EventNameList: []string{},
	}

	actual := example_func(&ppt)
	expected := 10 // something
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
