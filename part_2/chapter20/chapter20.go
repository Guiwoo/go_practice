package chapter20

import "fmt"

/**
- Interface 구체화된(Concrete Object => 구현이 있는 객체 를 의미)
- 객체 가 아닌 추상화된 상호작용으로 관계 를 표현

*/

type customer interface {
	string() string
}

type student struct {
	name   string
	age    int
	gender string
}

func (s student) string() string {
	return fmt.Sprintf("Hello My name is %v and %v years old", s.name, s.age)
}

type studentT struct {
	name string
}

func (s studentT) string() string {
	return fmt.Sprintf("Hello My name is %v and years old", s.name)
}

func Ex02() {
	s := student{}
	s.name = "1번"
	s2 := studentT{}
	s2.name = "2번"
	c := customer(s)
	c2 := customer(s2)

	arr := []customer{c, c2}
	fmt.Printf("This is type %T\n", arr)
	for _, v := range arr {
		fmt.Println(v)
	}
}

/**
fedex
*/

type Sender interface {
	Send(parcel string)
}

type FedexSender struct {
	//..
}

func (f *FedexSender) Send(parcel string) {
	fmt.Printf("Fedex sends one of %v parcel", parcel)
}

type PostSender struct {
	//..
}

func (p *PostSender) Send(parcel string) {
	fmt.Printf("우체국 sends one of %v parcel", parcel)
}

/**
 내부동작을 감춰서 서비스 제공자 와 사용자 모두에게 자유를 주는 방식을 추상화 (Abstraction)
Decoupling
*/

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

/**
DuckTyping
*/
