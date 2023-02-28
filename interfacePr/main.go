package main

import (
	"fmt"
)

type Sender interface {
	send()
}

type Email interface {
	sendEmail()
}
type email struct{}

func (e *email) sendEmail() {
	fmt.Println("고유 이메일")
}

type ExpressEmail struct {
	*email
	price int
}

func (e *ExpressEmail) send() {
	fmt.Println("Send Express Email")
}
func (e *ExpressEmail) sendEmail() {
	fmt.Println("holy?")
}

func main() {
	var a Sender = &ExpressEmail{
		&email{},
		500,
	}

	if v, ok := a.(*ExpressEmail); ok {
		v.send()
		v.email.sendEmail()
	}
}
