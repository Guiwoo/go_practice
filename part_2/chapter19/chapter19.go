package chapter19

import (
	"fmt"
	time "time"
)

/**
메서드 는 타입에 속한 함수
*/

type account struct {
	balance int
}

func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}
func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

func Ex01() {
	a := &account{100}

	withdrawFunc(a, 10)

	a.withdrawMethod(30)

}

// 포인트 타입 메서드 vs 값 타입 메서드
func (a1 *account) withdrawPointer(amount int) {
	a1.balance -= amount
}

func (a account) withdrawValue(amount int) {
	a.balance -= amount
}

// 언제 값을 쓸지 포인터 를 쓸지 이걸 타입에 따라 다르다
// 객체안의 데이터가 변경 되었을때 서로 다른 사람이냐
// go 에는 생성자 소멸자 가 없음

func Ex02() {
	var mainA *account = &account{100}
	mainA.withdrawPointer(30)
	fmt.Println(mainA.balance)

	mainA.withdrawValue(20)
	fmt.Println(mainA.balance)
}

type myInt int

func (m myInt) add(n int) int {
	return int(m) + n
}

/**
클래스 와 상속을 지원하지 않고 메서드 와 인터페이스 만 지
*/

type acc struct {
	balance   int
	firstName string
	lastName  string
}

func (a *acc) withdrawPointer(amount int) {
	a.balance -= amount
}

func (a acc) withdrawValue(amount int) {
	a.balance -= amount
}
func (a acc) withdrawReturnValue(amount int) acc {
	a.balance -= amount
	return a
}

func Ex03() {
	var mainA *acc = &acc{100, "Guiwoo", "Park"}
	mainA.withdrawPointer(10)
	fmt.Println(mainA)

	mainA.withdrawValue(19)
	fmt.Println(mainA.balance)

	var mainB acc = mainA.withdrawReturnValue(21)
	fmt.Println(mainB.balance)

	mainB.withdrawPointer(30)
	fmt.Println(mainB.balance)
}

type Courier struct {
	name string
}

type Product struct {
	name  string
	price int
	id    int
}

type Parcel struct {
	pdt           *Product
	shipTime      time.Time
	deliveredTime time.Time
}

func (c *Courier) SendProduct(pdt *Product) *Parcel {
	rst := &Parcel{}
	rst.pdt = pdt
	rst.shipTime = time.Now()
	return rst
}

func (p *Parcel) Delivered() *Product {
	p.deliveredTime = time.Now()
	return p.pdt
}
