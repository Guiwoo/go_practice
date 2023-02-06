package chapter13

import (
	"fmt"
	"unsafe"
)

type House struct {
	Address  string
	Size     int
	Price    float64
	Category string
}

func Ex01() {
	var house House
	house.Address = "서울"
	house.Size = 28
	house.Price = 100_000_000_000
	house.Category = "매매"

	fmt.Println(house)
}

type User struct {
	Name string
	Id   string
	Age  int
}

type VIPUser struct {
	UserInfo User
	VIPLevel int
	Price    int
}

func Ex04() {
	user := User{"Guiwoo", "1234", 31}
	vip := VIPUser{user, 2, 10}

	fmt.Println(user)
	fmt.Println(vip)
}

//embedded way 바로 접근가능하다 내장된 필드로 변경됨 즉 해당 타입에 귀속

func Ex05() {
	type user struct {
		num1 int8
		num2 int
		num3 int8
		num4 int
		num5 int8
	}
	// 8바이트 보다 작은 필드는 앞으로 모아서 배치하면 좋음
	usr := user{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(usr))
}

func Ex06() {
	type Product struct {
		Name  string
		Price int
		Speed float64
	}

}
