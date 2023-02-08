package chapter22

import (
	"container/list"
	"fmt"
)

func PlayWithList() {
	// linked list 네 bfs 해보자 이걸로 queue 니깐
	myList := list.New()
	for i := 0; i < 10; i++ {
		myList.PushBack(i)
	}

	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
