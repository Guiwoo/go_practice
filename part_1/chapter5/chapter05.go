package chapter5

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// fmt 를 이용한 표준 입출력 방법

func Ex01() {
	var a int = 10
	var b int = 20
	var f float64 = 327832942.1234

	// ln 은 저렇게 써도 , 사이에 공백을 넣어주지만 그냥 프린트는 다 붙여버림
	fmt.Print("a:", a, "b:", b)
	fmt.Println("a:", a, "b:", b, "f:", f)
	n, err := fmt.Println("a")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(n)
	fmt.Printf("a: %d b: %d f : %f\n", a, b, f)
}

func Ex02() {
	var name = "guiwoo"
	fmt.Fprintf(os.Stdout, "My name is : %v", name)
}

func Ex03() {
	var a = 123
	var b = 456
	var c = 123456789

	fmt.Printf("%5d, %5d\n", a, b)    // 최소너비 설정 후 우측정렬
	fmt.Printf("%05d, %05d\n", a, b)  //남는공간 0으로채움
	fmt.Printf("%-5d, %-05d\n", a, b) // 좌측정렬

	fmt.Printf("%5d, %5d\n", c, c)
	fmt.Printf("%05d, %05d\n", c, c)
}

func Ex04() {
	var a = 324.1346
	var b = 3.14

	fmt.Printf("%08.2f\n", a)
	fmt.Printf("%08.2g\n", a)
	fmt.Printf("%8.5g\n", a)
	fmt.Printf("%f\n", b)
}

func Ex05() {
	var a int
	var b int

	n, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}

//표준 입력스트림 은 fifo 구조로 Queue

func Ex06() {
	var a int
	var b int

	n, err := fmt.Scanf("%d %d", &a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(a, b)
	}
}

func Ex07() {
	stdin := bufio.NewReader(os.Stdin)

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}
	n, err = fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}
}

func Ex08() {
	var a = 123
	var b int = 4567
	f := 3.14159269

	fmt.Fprintf(os.Stdout, "%20d\n", a)
	fmt.Fprintf(os.Stdout, "%020d\n", b)
	fmt.Fprintf(os.Stdout, "%20.2f\n", f)
}
