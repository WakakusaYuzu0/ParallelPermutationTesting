package demo

import (
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
	expected := 10 /* 10 / ((3/1)-2) */
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

// passするテストケース
func TestDemo2(t *testing.T) {
	ppt := parallelpermutationtesting.Parapermtest{
		EventNameList: []string{
			"getThree", "getOne", "getTwo",
		},
	}

	actual := Demo_func(&ppt)
	expected := -10 /* 10 / ((2/1)-3) */
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

// divide by 0でfailするテストケース
func TestDemo3(t *testing.T) {
	ppt := parallelpermutationtesting.Parapermtest{
		EventNameList: []string{
			"getOne", "getTwo", "getThree",
		},
	}

	actual := Demo_func(&ppt)
	expected := 20 /* 10 / ((3/2)-1) */
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
