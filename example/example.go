package example

import (
	"fmt"
)

var ppt *Parapermtest

func getOne(ch chan int) {
	ppt.waitEvent("getOne")
	ch <- 1
	ppt.done()
}

func getTwo(ch chan int) {
	ppt.waitEvent("getTwo")
	ch <- 2
	ppt.done()
}

func getThree(ch chan int) {
	ppt.waitEvent("getThree")
	ch <- 3
	ppt.done()
}

func example_func(_ppt *Parapermtest) {
	ppt = _ppt

	ch := make(chan int, 3)

	go getOne(ppt, ch)
	go getTwo(ppt, ch)
	go getThree(ppt, ch)

	a := <-ch
	b := <-ch
	c := <-ch

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)

	fmt.Println("c/b - a:", c/b-a)
	fmt.Println("10 / (c/b - a)", 10/(c/b-a))
}
