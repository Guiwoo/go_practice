package chapter22

import (
	"fmt"
)

/**
맵 키:벨류 자료형 구조
*/
// 순서 보장을 안해주는 맵 입니다.

func Ex01() {
	m := make(map[string]int)

	m["a"] = 169
	m["b"] = 170
	m["c"] = 190

	for i, v := range m {
		fmt.Println(i, v)
	}
}

type Queue[T any] struct {
	value *[]T
}

func (q *Queue[T]) NewQueue() {
	q.value = &[]T{}
}

func (q *Queue[T]) Size() int {
	return len(*q.value)
}
func (q *Queue[T]) isEmpty() bool {
	return q.Size() == 0
}
func (q *Queue[T]) Add(a T) bool {
	*q.value = append(*q.value, a)
	// q.value = &append(*q.value,a)
	return true
}
func (q *Queue[T]) Pop() T {
	if q.isEmpty() {
		panic("This Queue is Empty")
	}
	arr := *q.value
	*q.value = arr[1:]

	// arr := *q.value
	// nArr := arr[1:]
	// q.value = &nArr

	// arr = *q.value
	// q.value = &arr[1:]
	return arr[0]
}

func PlayWithQueue() {
	q := Queue[int]{}
	q.NewQueue()

	q.Add(1)
	q.Add(2)
	q.Add(3)
	q.Add(4)

	fmt.Println(*q.value)

	for !q.isEmpty() {
		a := q.Pop()
		fmt.Print(a, " ")
	}
}
