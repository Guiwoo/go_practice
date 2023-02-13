package chapter27

import "fmt"

/**
Solid 란 Oop

단일책임 원칙 SRP
개방폐쇄 원칙 OCP
리스코프 치환 원칙 LSP
인터페이스 분리원칙 ISP
의존관계 역전 원칙 DIP


GO 도 객체 중심 프로그래밍 이 가능하다.
*/

/**
SRP 위반 사례 보고서 와 전송을 분리해서 만들어야 한다.
보고서 인터페이스 => 전송 담당
*/

type FinanceReport struct {
	Sibal string
}

// SendReport 전송 로직
func (r *FinanceReport) SendReport(email string) {
	// ... Logic
}

/**
변경후
*/

type Report interface {
	Report() string
}

func (r *FinanceReport) Report() string {
	return r.Sibal
}

type ReportSender struct {
}

func (r *ReportSender) SendReport(report Report) {
	//
	report.Report()
}

/*
*
개방폐쇄 원칙
확장에 열려있고 ,변경에는 닫혀있는경우
*/
type SendType interface {
	Send()
}
type Email struct{}

func (e *Email) Send() {
	fmt.Println("Email Sender")
}

type Fex struct{}

func (f *Fex) Send() {
	fmt.Println("Fex Sender")
}

func SendReport(r *Report, method SendType, receiver string) {
	switch method {
	case method.(*Email):
		fmt.Println("Case Sender")
	case method.(*Fex):
		fmt.Println("Case Fex")
	default:
		fmt.Println("None")
	}
}

/**
리스코프 치환 원칙 lsp
q(x) 를 타입 T 의 객체 x 에 대해 증명할수있는 속성이라하자.
S 가 T 의 하위 타입이라면 S 의 객체 y 에 대해 증명할수 있어야한다.
*/

/**
ISP 인터페이스 분리원칙
*/

/**
의존관계 역전원칙
1. 상위모듈은 하위모듈에 의존해서는 안된다 둘다 추상모듈에 의존해야한다.
2. 구체화된 모듈은 추상모듈에 의존해야한다.
*/
