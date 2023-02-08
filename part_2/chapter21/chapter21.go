package chapter21

import (
	"fmt"
	"os"
)

/**
함수
*/

func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func deferTest() {
	f, err := os.Create("Text.txt")
	if err != nil {
		fmt.Println("Failed to Create File")
		return
	}
	defer fmt.Println("호출됩니다")
	defer f.Close()

	fmt.Println("Write a hello on Text.txt file")
	fmt.Fprintf(f, "Hello World %v", "It's me")
}

func add(a, b int) int {
	return a + b
}
func mult(a, b int) int {
	return a * b
}

type opFunc func(int, int) int

func getOp(op string) opFunc {
	switch op {
	case "+":
		return add
	case "*":
		return mult
	default:
		return nil
	}
}

func getOp2(op string) opFunc {
	if op == "+" {
		return func(a, b int) int {
			return a + b*a
		}
	} else if op == "*" {
		return func(a, b int) int {
			return a * b * a * b
		}
	} else {
		return nil
	}
}

func CaptureLoop() {
	f := make([]func(), 3)
	fmt.Println("Print line")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}
	for l := 0; l < 3; l++ {
		f[l]()
	}
}

func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("Print line 2")
	for i := 0; i < 3; i++ {
		v := i
		f[i] = func() {
			fmt.Println(v)
		}
	}
	for l := 0; l < 3; l++ {
		f[l]()
	}

}

type Writer func(string)

func writeHello(writer Writer) {
	writer("Hello World")
}
func writeMain() {
	f, err := os.Create("Test.txt")
	if err != nil {
		fmt.Println("Failed to create a file !")
		return
	}
	defer f.Close()

	writeHello(func(msg string) {
		fmt.Fprintln(f, msg)
	})
}
