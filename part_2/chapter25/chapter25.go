package chapter25

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/**
채널 과 컨텍스트
채널 이란 ? 고 루틴 간의 의사소통 수단 메세지 큐 fifo 자료구조
var messages chan string = make(chan string)
message <- "This is a message" 채널 안으로 데이터가 들어간다 ~
var msg string = <- messages 채널에서 뺀다 ~
*/

func Ex01() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)
	ch <- 9
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch

	time.Sleep(time.Second)
	fmt.Println("Square : ", n*n)
	wg.Done()
}

func Ex02() {
	var wg = &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(1)
	go sq2(wg, ch)
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	close(ch)
	wg.Wait()
	fmt.Println("Never print")
}

func sq2(wg *sync.WaitGroup, ch chan int) {
	for n := range ch {
		fmt.Println("Square :", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}

// select 문 switch 와 유사 하지만 channel 을 사용할떄 씀 여러 채널에서 동시에 데이터 가져올떄 사용
// 단 하나의 케이스가 종료되면 끝남 그래서 주로 for 를 앞에걸어 무한 루프 돌림

func Ex03() {
	var wg = &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(1)
	go sq3(wg, ch)
	for i := 0; i < 10; i++ {
		ch <- i * i * i
	}
	wg.Wait()
}

func sq3(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)
	terminate := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-terminate:
			fmt.Println("Terminated Tok")
			wg.Done()
			return
		case n := <-ch:
			fmt.Println("Square: ", n*n)
			time.Sleep(time.Second)
		}
	}
}

type Car struct {
	Body   string
	Tire   string
	Engine string
}

func Ex04() {
	var wg = &sync.WaitGroup{}
	var stTime = time.Now()
	bodyChan := make(chan *Car)
	tireChan := make(chan *Car)

	wg.Add(3)
	go createEngine(bodyChan, wg)
	go createBody(bodyChan, tireChan, wg)
	go createTire(tireChan, wg, stTime)

	wg.Wait()
}

func createEngine(bodyChan chan *Car, wg *sync.WaitGroup) {
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)
	for {
		select {
		case <-tick:
			car := &Car{Engine: "gasoline"}
			fmt.Println("⚙️ Create Engine")
			bodyChan <- car
		case <-after:
			close(bodyChan)
			wg.Done()
			return
		}
	}
}

func createBody(bodyChan, tireChan chan *Car, wg *sync.WaitGroup) {
	for car := range bodyChan {
		time.Sleep(time.Second)
		fmt.Println("🚘 Create Body")
		car.Body = "Carbonate frame"
		tireChan <- car
	}
	wg.Done()
	close(tireChan)
}

func createTire(tireChan chan *Car, wg *sync.WaitGroup, time2 time.Time) {
	for car := range tireChan {
		time.Sleep(time.Second)
		fmt.Println("🛞 Installing Tire")
		car.Tire = "Continental Summer Tire"

		dur := time.Now().Sub(time2)
		fmt.Printf("✅ Complete Car Duration is %.2f\n", dur.Seconds())
	}
	wg.Done()
}

// 문맥

//var wg sync.WaitGroup

func Ex05() {

	var wg = &sync.WaitGroup{}
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())

	go printEverySecond(ctx, wg)
	time.Sleep(5 * time.Second)
	cancel()

	wg.Wait()
}

func printEverySecond(ctx context.Context, wg *sync.WaitGroup) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		case <-tick:
			fmt.Printf("Tick Tick\n")
		}
	}
}

/**
발행 구독 패턴 pub/sub pattern similar as observer pattern
*/

func Ex06() {
	ctx, cancel := context.WithCancel(context.Background())

	Wg.Add(4)
	publisher := NewPublisher(ctx)
	subscriber1 := NewSubscriber("Guiwoo", ctx)
	subscriber2 := NewSubscriber("Park", ctx)

	go publisher.Update()

	subscriber1.Subscribe(publisher)
	subscriber2.Subscribe(publisher)

	go subscriber1.Update()
	go subscriber2.Update()

	go func() {
		tick := time.Tick(time.Second * 2)
		for {
			select {
			case <-tick:
				publisher.Publish("Hello EveryOne")
			case <-ctx.Done():
				Wg.Done()
				return
			}
		}
	}()

	fmt.Scanln()
	cancel()

	Wg.Wait()
}
