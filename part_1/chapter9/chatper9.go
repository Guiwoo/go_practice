package chapter9

import (
	"fmt"
	"time"
)

func Ex01(age int) {
	if age >= 10 && age <= 15 {
		fmt.Println("You are young!")
	} else if age > 30 || age < 20 {
		fmt.Println("you are not 20s")
	} else {
		fmt.Println("Best age of your life")
	}
}

var cnt int = 0

// short circuit
func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}

func getMyage() (int, bool) {
	return 33, true
}

func ageProcess() {
	if age, ok := getMyage(); ok && age < 20 {
		fmt.Println("You ar young", age)
	} else if ok && age < 30 {
		fmt.Println("Nice age ", age)
	} else if ok {
		fmt.Println("YOu are beautiful", age)
	} else {
		fmt.Println("error")
	}
	//fmt.Println("Your age is",age) age 는 if 스코프 안에 있기때문에 if
	// 가 끝난 바깥 스코프에서 접근 불가
}
func getTemp(day int) (temp int, rain float32) {
	//temp function
	temp = 20
	//get rain percentage
	rain = 80
	return
}
func GetWeather() {
	day := time.Now().Day()
	temp, rain := getTemp(day)
	if temp < 20 && rain < 80 {
		fmt.Println("Clumsy day")
	} else if temp < 30 && rain < 70 {
		fmt.Println("So Hot")
	} else {
		fmt.Println("It's insane")
	}
}
