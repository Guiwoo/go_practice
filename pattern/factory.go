package main

import "fmt"

type Camera interface {
	CreateCamera()
	CreateChip()
}

type PhoneFactory interface {
	create() Camera
}
type Phone struct {
	phone PhoneFactory
}
type PhoneType uint

const (
	GALAXY PhoneType = iota
	IPHONE
	HUAWEI
	PIXEL
)

func (p *Phone) newPhone(pht PhoneType) {
	switch pht {
	case GALAXY:
		fmt.Println("Galaxy")
	case IPHONE:
		fmt.Println("Iphone")
	case HUAWEI:
		fmt.Println("Huawei")
	case PIXEL:
		fmt.Println("Pixel")
	}
}

func main() {
	var a Phone
	a.newPhone(GALAXY)
}
