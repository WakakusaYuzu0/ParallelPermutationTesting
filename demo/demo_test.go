package demo

import (
	"fmt"
	"testing"
	"testtool/parallelpermutationtesting"
)

func TestDemo(t *testing.T) {
	ppt := parallelpermutationtesting.Parapermtest{
		// IsDryRun:      true,
		EventNameList: []string{
			"getTwo", "getOne", "getThree",
		},
	}

	actual := Demo_func(&ppt)
	fmt.Println(actual)
	expected := 10 // something
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

func TestDemo2(t *testing.T) {
	ppt := parallelpermutationtesting.Parapermtest{
		// IsDryRun:      true,
		EventNameList: []string{
			"getTwo", "getThree", "getOne",
		},
	}

	actual := Demo_func(&ppt)
	fmt.Println(actual)
	expected := -5 // something
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
