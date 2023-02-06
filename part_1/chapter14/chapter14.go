package chapter14

import (
	"awesomeProject/part_1/chapter13"
	"fmt"
)

// Pointer
// 포인터 로 연산 정도 가능 하다 ,단 c 나 c++ 보다 는 덜하다
/**
원리 포인터 는 메모리 주소를 값으로 갖는 타입
*int => int 타입 의 메모리 주소 값
여러 포인트 변수 가 하나의 변수를 가르킬수 있다.
*/

func Ex01() {
	var a int = 500
	var p *int

	p = &a
	fmt.Printf("p 의 값은 : %p\n", p)
	fmt.Printf("메모리 값 %d\n", *p)
}

// pointer 에도 == 연산자 사용 이 가능하다
// pointer 기본값 은 nil == null

func Ex02() {
	var a int = 10
	var b int = 20

	var p1 *int = &a
	var p2 *int = &a
	var p3 *int = &b
	fmt.Printf("p1 == p2 :%v \n", p1 == p2)
	fmt.Printf("p2 == p3 :%v \n", p2 == p3)
}

// 함수 의 인자 ? rValue 값을 복사 함
// java ? value => Reference 값 이 들어오지만, go => 복사가 일어남

func Ex03() {
	type Data struct {
		value int
		data  [200]int
	}
	changeData := func(arg *Data) {
		arg.value = 999
		arg.data[100] = 999
	}

	var data Data
	changeData(&data)
	fmt.Printf("value =%d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])
}

// Instance 인스턴스 는 메모리 할당된 데이터 실체 포인터 가 가르키는 데이터
// Instance 는 사용되어 지지 않을때 가비지 컬렉터에서 수거함 수거하는 시점이 언제일까
// new(Data) 이렇게 하면 포인터 타입의 값을 반환 함

func Ex06() {
	u := &chapter13.User{}
	u.Age = 30
	fmt.Println(u)
}

// stack && heap
func Ex07() {
	type user struct {
		Name string
		Age  int
	}

	newUser := func(name string, age int) *user {
		var u = user{name, age}
		return &u
	}

	userPointer := newUser("AAA", 23)
	fmt.Println(userPointer)
}
