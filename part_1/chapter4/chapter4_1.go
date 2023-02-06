package main

import "fmt"

// 변수 는 값을 저장하는 메모리 공간

func Ex01() {
	var a int = 10
	var msg string = "Hello Variable"

	a = 20
	msg = "Good morning"
	fmt.Println(msg, a)
}

func Ex02() {
	var a int = 3
	var b int
	var c = 4
	d := 5

	fmt.Println(a, b, c, d)
}

// 실수 타입에서 정수타입 으로 가면 소수점 삭제
// 큰범위 에서 작은범위 타입 으로 변환하면 값이 달라질수 있다.

func Ex03() {
	a := 3
	var b float64 = 3.5
	var c int = int(b)

	var e int64 = 7
	f := int64(c) * e

	fmt.Println(a, f)
}

// 값을 초과하는 범위 상위 1바이트 가 없어지기 때문에 최소값 리턴

func Ex04() {
	var a int16 = 3456   // 3456
	var c int8 = int8(a) // -128

	fmt.Printf("%v %v", a, c)
}

// go 의 변수 스코프
// 글로벌 과 함수 스코프 에 동일한 변수 선언시 가장 가까운 스코프의 변수 먼저 사용

var myvariable1 int = 100

func Ex05() {
	// local variables inside the main function
	// it is same as global variable
	var myvariable1 int = 200

	// Display the value
	fmt.Printf("The value of myvariable1 is : %d\n",
		myvariable1)
}

func Ex06() {
	var a float32 = 1234.523
	var b float32 = 3456.123
	var c float32 = a * b
	var d float32 = c * 3
	// 금융 회계 에서는 이런 실수 계산 할때 주의를 해야한다.
	// 약 100 이상의 큰차이가 생기는 오류가 발생
	fmt.Printf("a : %v \n b : %v \n c : %v \n d: %v", a, b, c, d)
}

func Ex07() {
	a := 3
	var b = 3.1415
	c := "Hello World"
	d := 'H'
	e := int32(10)
	var f float32 = 3.1415

	fmt.Printf("Type of a is : %T \n Type of b is : %T", a, b)
	fmt.Printf("Type of c is : %T \n Type of d is : %T", c, d)
	fmt.Printf("Type of e is : %T \n Type of b is : %T", e, f)
}
