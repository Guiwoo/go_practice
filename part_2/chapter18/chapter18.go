package chapter18

import (
	"fmt"
	"sort"
)

/**
일반 Go 배열 은 정적 사이즈 를 가지고 있음 [3] 한번 정한 사이즈는 변동 불가능한 문제점 을 야기함
*/

func bookEx01() {
	var slice []int
	if len(slice) < 1 {
		fmt.Println("slice is Empty", slice)
	}
	slice[1] = 10 // 에러 발생 슬라이스 길이 이상을 호출하려고 했기 때문에 index out of range panic error
	fmt.Println(slice, len(slice))
}

func BookEx02() {
	var array = [...]int{1, 2, 3}
	var slice = []int{1, 2, 3}
	var slice2 = make([]int, 3)
	fmt.Printf("Declare as [...] is %T\n", array)
	fmt.Printf("Declare as slice is %T\n", slice)
	fmt.Printf("Declare as slice is %T", slice2)
}

func BookEx03() {
	var slice = []int{1, 2, 3}

	for i := 0; i < len(slice); i++ {
		slice[i] *= 2
	}
	for _, v := range slice {
		fmt.Print(v, " ")
	}
}

func BookEx04() {
	var slice = []int{1, 2, 3}
	slice2 := append(slice, 4)
	slice3 := append([]byte("hello"), "world"...)

	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(string(slice3))
}

func BookEx05() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5, 6)

	fmt.Println("slice 1 is", slice1, len(slice1), cap(slice1))
	fmt.Println("slice 2 is ", slice2, len(slice2), cap(slice2))

	slice1[1] = 100

	fmt.Println("After change second value")
	fmt.Println("slice 1 is", slice1, len(slice1), cap(slice1))
	fmt.Println("slice 2 is ", slice2, len(slice2), cap(slice2))

	slice1 = append(slice1, 500)
	fmt.Println("After append 500")
	fmt.Println("slice 1 is", slice1, len(slice1), cap(slice1))
	fmt.Println("slice 2 is ", slice2, len(slice2), cap(slice2))

}

type student struct {
	name  string
	age   int
	score int
	rate  float32
}
type students []student

func (s students) Len() int           { return len(s) }
func (s students) Less(i, j int) bool { return s[i].score < s[j].score }
func (s students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func BookEx06() {
	arr := []student{
		{"나통키", 13, 45, 78.4},
		{"오맹태", 16, 24, 67.4},
		{"오동도", 18, 54, 50.8},
		{"황금산", 16, 36, 89.7},
	}

	sort.Sort(students(arr))
}

func Ex01() {
	addNum := func(slice []int) {
		slice = append(slice, 4)
	}

	slice := []int{1, 2, 3}
	addNum(slice)
	fmt.Println(slice)
}
