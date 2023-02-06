package chapter10

import "fmt"

/*
*
Switch
*/
func Ex01(a int) {
	switch a {
	case 1:
		fmt.Println("Bull shit")
	case 2:
		fmt.Println("Holy Moly")
	}
}
func getMyAge() int {
	return 31
}

func Ex02() {
	switch age := getMyAge(); age {
	case 10:
		fmt.Println("Your are young")
	case 20:
		fmt.Println("Pair 3")
	default:
		fmt.Println("My age is ", age)
	}
}

type Direction int

const (
	None Direction = iota
	North
	East
	South
	West
)

func getDirection(angle float64) Direction {
	switch {
	case angle < 90:
		return North
	case angle < 180:
		return West
	case angle < 270:
		return East
	case angle < 360:
		return South
	default:
		return None
	}
}