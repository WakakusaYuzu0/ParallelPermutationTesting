package demo

import (
	"fmt"
	"testtool/parallelpermutationtesting"
)

// 並行処理テストオブジェクトを定義
var ppt *(parallelpermutationtesting.Parapermtest)

// テストしたい関数
func Demo_func(_ppt *parallelpermutationtesting.Parapermtest) int {
	// 並行処理テストのインスタンスを受け取る
	ppt = _ppt

	ch := make(chan int, 3)

	go getOne(ch)
	go getTwo(ch)
	go getThree(ch)

	a := <-ch
	b := <-ch
	c := <-ch

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)

	fmt.Println("c/b - a:", c/b-a)
	fmt.Println("10 / (c/b - a)", 10/(c/b-a))

	return 10 / (c/b - a)
}

// goroutineで呼ぶ関数
func getOne(ch chan int) {
	// goroutineの処理単位をWaitEventとDoneで区切る
	ppt.WaitEvent("getOne")
	ch <- 1
	ppt.Done()
}

// goroutineで呼ぶ関数
func getTwo(ch chan int) {
	ppt.WaitEvent("getTwo")
	ch <- 2
	ppt.Done()
}

// goroutineで呼ぶ関数
func getThree(ch chan int) {
	ppt.WaitEvent("getThree")
	ch <- 3
	ppt.Done()
}
