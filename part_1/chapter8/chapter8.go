package chapter8

import "fmt"

func Ex01() {
	const pi float64 = 3.141592653589793238
	var PI2 float64 = 3.141592653589793238

	PI2 = 4

	fmt.Println(pi)
	fmt.Println(PI2)
}

const Pig int = 0
const Cow int = 1
const Chicken int = 2

func Ex02(animal int) {
	if animal == Pig {
		fmt.Println("꿀꿀")
	} else if animal == Cow {
		fmt.Println("음머")
	} else if animal == Chicken {
		fmt.Println("꼬끼어")
	} else {
		fmt.Println("Don't know anything")
	}
}

const (
	Red   int = iota
	Blue  int = iota
	Green int = iota
)

const (
	C1 uint = iota + 1
	C2
	C3
)
