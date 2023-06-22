// package parallel_perm_test
package main

// import (
// 	"permutation"
// )
import (
	"fmt"
	"time"
)

type Parapermtest struct {
	isInspection bool
	eventIdList  []int
}

// イベントを登録する
func (ppt *Parapermtest) addEvent() {
	fmt.Println("add event")
	fmt.Println(1)
	if len(ppt.eventIdList) == 0 {
		ppt.eventIdList = []int{0}
	} else {
		ppt.eventIdList = append(ppt.eventIdList, ppt.eventIdList[len(ppt.eventIdList)-1]+1)
	}

	return
}

// イベントを実行する
func nanika() {

}

func (ppt *Parapermtest) done() {
	ppt.eventIdList = ppt.eventIdList[1:]
	time.Sleep(1 * time.Second)
}

func (ppt *Parapermtest) waitEvent() int {
	myid := 0
	if ppt.isInspection {
		ppt.addEvent()
	} else {
		/*for (<- eventChannel) != myid {
			time.Sleep(1 * time.Second)
		}*/
		if len(ppt.eventIdList) == 0 {
			return myid
		}
		for ppt.eventIdList[0] != myid {
			time.Sleep(1 * time.Second)
		}
	}
	return myid
}

func (ppt *Parapermtest) waitEvent2() int {
	myid := 1
	if ppt.isInspection {
		ppt.addEvent()
	} else {
		/*for (<- eventChannel) != myid {
			time.Sleep(1 * time.Second)
		}*/
		if len(ppt.eventIdList) == 0 {
			return myid
		}
		for ppt.eventIdList[0] != myid {
			time.Sleep(1 * time.Second)
		}
	}
	return myid
}

func (ppt *Parapermtest) waitEvent3() int {
	myid := 2
	if ppt.isInspection {
		ppt.addEvent()
	} else {
		/*for (<- eventChannel) != myid {
			time.Sleep(1 * time.Second)
		}*/
		if len(ppt.eventIdList) == 0 {
			return myid
		}
		for ppt.eventIdList[0] != myid {
			time.Sleep(1 * time.Second)
		}
	}
	return myid
}

// 順列を求める
// func getPermList(events []int) [][]int {
// 	return Permute(events)
// }

/*
func doWholeTest() {
	ppt := Parapermtest{
		isInspection: true,
		eventIdList:  []int{},
	}
	// eventを登録しようね

	// イベントで順列を作るよ

	// 作った順列を一つずつ実行するよ
	for _, perm := range getPermList() {
		doTest(&ppt, perm)
	}
	// エラーが出た！どうすんの！？
}
*/

func doTest() {
	ppt := Parapermtest{
		isInspection: true,
		eventIdList:  []int{},
	}
	// dryRun(&ppt)
	run(&ppt)

}

func run(ppt *Parapermtest) {
	permList := getPermList([]int{0, 1, 2})
	fmt.Println(permList)
	for _, perm := range permList {
		ppt := Parapermtest{
			isInspection: false,
			eventIdList:  perm,
		}

		fmt.Println("=======================")
		demo_func(&ppt)
		time.Sleep(5 * time.Second)
	}
}
