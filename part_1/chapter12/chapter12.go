package chapter12

import "fmt"

/**
Array
배열의 선언 및 초기화
var nums [5]int
*/

func Ex01() {
	var t [5]float64 = [5]float64{24.0, 25.9, 21.2, 13.3, 11.2}
	for i := 0; i < len(t); i++ {
		fmt.Println(t[i])
	}
	for idx, val := range t {
		fmt.Println(idx, val)
	}
}

func Ex02() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{500, 400, 300, 200, 100}

	for i, v := range a {
		fmt.Printf("a[%d] = %d \n", i, v)
	}
	fmt.Println()
	for i, v := range b {
		fmt.Printf("b[%d] = %d \n", i, v)
	}
	b = a

	for i, v := range b {
		fmt.Printf("b[%d] = %d \n", i, v)
	}
}

func Ex03() {
	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{5, 6, 7, 8, 9},
	}

	for _, err := range a {
		for _, v := range err {
			fmt.Printf("%d ", v)
		}
	}
}