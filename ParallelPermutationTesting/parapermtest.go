// package parallel_perm_test
package parallelpermutationtesting

// import (
// 	"permutation"
// )
import (
	"fmt"
	"time"
)

type Parapermtest struct {
	IsDryRun bool
	// eventIdList []int
	EventNameList []string
	// eventIdName map[int]string
	// eventNameId map[string]int
}

func (ppt *Parapermtest) WaitEvent(eventName string) bool {
	if ppt.IsDryRun {
		ppt.addEvent(eventName)
	} else {
		/*for (<- eventChannel) != myid {
			time.Sleep(1 * time.Second)
		}*/
		if len(ppt.EventNameList) == 0 {
			return true
		}
		for ppt.EventNameList[0] != eventName {
			time.Sleep(1 * time.Second)
		}
	}
	return true
}

// イベントを登録する
func (ppt *Parapermtest) addEvent(eventName string) {
	fmt.Println("add event")
	// if len(ppt.eventIdList) == 0 {
	// 	ppt.eventIdList = []int{0}
	// } else {
	// 	ppt.eventIdList = append(ppt.eventIdList, ppt.eventIdList[len(ppt.eventIdList)-1]+1)
	// }
	ppt.EventNameList = append(ppt.EventNameList, eventName)

	return
}

func (ppt *Parapermtest) Done() {
	if !ppt.IsDryRun {
		ppt.EventNameList = ppt.EventNameList[1:]
		time.Sleep(1 * time.Second)
	}
}

func doTest() {
	ppt := Parapermtest{
		IsDryRun:      true,
		EventNameList: []string{},
	}
	dryRun(&ppt)
	// run(&ppt)

}

func dryRun(ppt *Parapermtest) {
	// demo.Demo_func(ppt)
	ppt.IsDryRun = false
	return
}

/*func run(ppt *Parapermtest) {
	ppt.runAllPermTest()
}

func (ppt *Parapermtest) runAllPermTest() {
	eventNameId := map[string]int{}
	eventIdName := map[int]string{}

	// nameListLen := len(ppt.EventNameList)
	idList := []int{}
	fmt.Println(ppt.EventNameList)
	for i, name := range ppt.EventNameList {
		fmt.Println(i, name)
		idList = append(idList, i)
		eventNameId[name] = i
		eventIdName[i] = name
	}
	fmt.Println(idList)

	permIdList := getPermList(idList)
	fmt.Println(permIdList)

	for _, perm := range permIdList {
		permNameList := []string{}
		fmt.Println(permNameList)
		for _, id := range perm {
			permNameList = append(permNameList, eventIdName[id])
		}

		ppt.EventNameList = permNameList

		fmt.Println("=======================")
		demo.Demo_func(ppt)
		time.Sleep(5 * time.Second)
	}
}
*/

/*
func (ppt *Parapermtest) waitEvent2() int {
	myid := 1
	if ppt.IsDryRun {
		ppt.addEvent()
	} else {
		for (<- eventChannel) != myid {
			time.Sleep(1 * time.Second)
		}
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
	if ppt.IsDryRun {
		ppt.addEvent()
	} else {
		for (<- eventChannel) != myid {
			time.Sleep(1 * time.Second)
		}
		if len(ppt.eventIdList) == 0 {
			return myid
		}
		for ppt.eventIdList[0] != myid {
			time.Sleep(1 * time.Second)
		}
	}
	return myid
}
*/

// 順列を求める
// func getPermList(events []int) [][]int {
// 	return Permute(events)
// }

/*
func doWholeTest() {
	ppt := Parapermtest{
		IsDryRun: true,
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
