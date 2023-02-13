package practice

import (
	"fmt"
	"strconv"
	"time"
)

func Ex01() {
	ch := make(chan string)

	for i := 0; i < 10; i++ {
		go func(i int) {
			for j := 0; j < 10; j++ {
				ch <- "Goroutine : " + strconv.Itoa(i)
			}
		}(i)
	}

	for cnt := 0; cnt < 100; cnt++ {
		fmt.Println(cnt, <-ch)
	}
	fmt.Println("Program is finished")
}

func receiver(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
func generator() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i * i
		}
		close(c)
	}()
	return c
}

func Ex02() {
	c := generator()
	receiver(c)
}

func fibo(ch chan int, quit chan bool) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			fmt.Println("X 콜")
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit Sign")
			return
		}
	}
}

func Ex03() {
	ch := make(chan int)
	quit := make(chan bool)
	n := 10

	go func(n int) {
		for i := 0; i < n; i++ {
			fmt.Println(<-ch)
		}
		quit <- false
	}(n)
	fibo(ch, quit)
}

func Announce(message string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println(message)
	}() // Note the parentheses - must call the function.
}

func F(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
