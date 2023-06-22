package main

import (
	"fmt"
)

func getOne(ppt *Parapermtest, ch chan int) {
	ppt.waitEvent()
	ch <- 1
	ppt.done()
}

func getTwo(ppt *Parapermtest, ch chan int) {
	ppt.waitEvent2()
	ch <- 2
	ppt.done()
}

func getThree(ppt *Parapermtest, ch chan int) {
	ppt.waitEvent3()
	ch <- 3
	ppt.done()
}

func demo_func(ppt *Parapermtest) {
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

	fmt.Println("a/b - c:", a/b-c)

	fmt.Println("10 / (a/b - c)", 10/(a/b-c))
}
