package chapter24

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
*
Go Routine
쓰레드 => 프로세스 로 부터 자원을 할당 받아 돌아가는것 으로 알고 있는데 ?
코어가 빠르게 쓰레드 를 교체한다. 마치 프로그램 이 동시에 실행되는 것처럼 보이는 기법 => 결국 코드는 하나씩 만 실행됨
cpu 는 가만히 있지만 os 에서 다음 실행 명령어를 넣어줌

단점 : 컨텍스트 스위칭 이 발생함
쓰레드 전환 에는 비용이 발생하는데 이걸 성능 상 문제 라고 함.
멀티 프로세스 ? => cpu 로 부터 자원을 할당받아 실행 되고 각 프로세스 간에는 특별한 통신수단을 통해 정보를 교환 하는것
멀티 쓰레드  ? => 프로세스 로 부터 스택,힙 영역을 할당 받아 각 스레드 별로 자원을 공유하는

고루틴 경량 쓰레드 왜 경량 이라는 호칭이 붙을끼 ? 자바 쓰레드랑 차이점이 무엇이 있을라나 검색해볼것
*/
func printNumber(sec time.Duration) {

	han := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n'}

	for _, v := range han {
		time.Sleep(sec * time.Millisecond)
		fmt.Printf("%c ", v)
	}
}
func Ex01() {
	var a int64 = 300
	var b int64 = 400

	go printNumber(time.Duration(a))
	go printNumber(time.Duration(b))
	time.Sleep(8 * time.Second)
}

// 서브 고루틴이 종료 될때 까지 대기 하는법
// sync.WaitGroup
var wg sync.WaitGroup

func sumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("from %d to %d sum is %d\n", a, b, sum)
	wg.Done()
}

/**
왜 좋은가 ? light weight thread
goRoutine != thread
시스템콜 , 네트워크 패킷 읽기 쓰기
이렇게 되면 아무리 생성해도 컨텍스트 스위칭 비용이 현저하게 저렴하다 Os 의 스위칭 보다
왜 ? 쓰레드가 코어의 갯수에 맞춰서만 생성 되기 떄문
*/

func Ex02() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go sumAtoB(1, 1000000)
	}
	wg.Wait()
}

// 동일 메모리자원을 여러 고루틴에서 접근할떄 문제 발생 ? 데드락,
// 이렇게 되면 ? 성능향상을 얻을수 없다. 왜 ? 1개만 접근하도록 제한 하기 떄문에
var mutex sync.Mutex

type Account struct {
	Balance int
}

func Ex03() {
	var wg sync.WaitGroup

	account := &Account{10}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				depositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func depositAndWithdraw(account *Account) {
	mutex.Lock()
	defer mutex.Unlock()

	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}

func Ex04() {
	rand.Seed(time.Now().UnixNano())

	wg.Add(2)
	fork := &sync.Mutex{}
	spoon := &sync.Mutex{}

	go dProblem("A", fork, spoon, "포크", "spoon")
	go dProblem("B", spoon, fork, "spoon", "fork")

	wg.Wait()
}

// 코드 안에서는 호락호락 하지 않다.
// 항상 발생 하지 않을수도 있고, 데드락이 터짐 간헐적 으로 그리고 파악하기도 힘듬

func dProblem(name string, first, second *sync.Mutex, firstName, lastName string) {
	for i := 0; i < 100; i++ {

		fmt.Printf("%s 밥 냠냠\n", name)
		first.Lock()
		fmt.Printf("%s %s gain\n", name, firstName)
		second.Lock()
		fmt.Printf("%s %s gain\n", name, lastName)

		fmt.Printf("%s eat food\n", name)

		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		second.Unlock()
		first.Unlock()
	}
}

// 동시성 해결방법 영역을 나누는 방법
type Job interface {
	Do()
}

type SquareJob struct {
	index int
}

func (j *SquareJob) Do() {
	fmt.Printf("%d start working\n", j.index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d finish work result : %d\n", j.index, j.index*j.index)
}

func Ex05() {
	var jobList [10]Job

	for i := 0; i < 10; i++ {
		jobList[i] = &SquareJob{i}
	}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		job := jobList[i]
		go func() {
			job.Do()
			wg.Done()
		}()
	}
	wg.Wait()
}
