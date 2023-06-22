package main

import (
	"fmt"
)

func homehome(ppt *Parapermtest) {
	ppt.waitEvent()
	fmt.Println("You are such a TENSAI!! 1")
	ppt.done()
}

func homehome2(ppt *Parapermtest) {
	ppt.waitEvent2()
	fmt.Println("You are such a TENSAI!! 2")
	ppt.done()
}

func homehome3(ppt *Parapermtest) {
	ppt.waitEvent3()
	fmt.Println("You are such a TENSAI!! 3")
	ppt.done()
}

func sample_func(ppt *Parapermtest) {
	go homehome(ppt)
	go homehome2(ppt)
	go homehome3(ppt)

	for len(ppt.eventIdList) > 0 {

	}
	return
}
