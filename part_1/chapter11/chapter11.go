package chapter11

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func Ex01() {
	var a int
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		a++
	}
	fmt.Println(a)
}

func Ex02() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("입력하세요 숫자를")
		var number int
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println("숫자를 입력하세요 . ")

			stdin.ReadString('\n')
			continue
		}
		fmt.Printf("입력하신 숫자는 %d 입니다.", number)
		if number%2 == 0 {
			break
		}
	}
	fmt.Println("For 문이 종료되었습니다.")
}

func Ex03() {
	a := 1
	b := 1

OuterFor:
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				break OuterFor
			}
		}
	}

	fmt.Printf("%d * %d = %d \n", a, b, a*b)
}
func find45(a int) (int, bool) {
	for b := 1; b <= 9; b++ {
		if a*b == 45 {
			return b, true
		}
	}
	return 0, false
}

func Ex04() {
	a := 1
	b := 0

	for ; a <= 9; a++ {
		var found bool
		if b, found = find45(a); found {
			break
		}
	}
	fmt.Printf("%d * %d = %d", a, b, a*b)
}

func Ex05() {
	var a int = 5
	for i := 0; i < 5; i++ {
		fmt.Println(strings.Repeat("*", a-i))
	}
}
