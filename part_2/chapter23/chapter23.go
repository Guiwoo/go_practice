package chapter23

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
오류 발생 에 따른 핸들링
에러타입 은 그냥 인터페이스에 에러 스트링으로 되어 있음
*/

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	line, _ := rd.ReadString('\n') //eof => end of file
	return line, nil
}

func WriteFile(filename string, line string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, line)
	return err
}

func Ex01() {
	const filename = "test.txt"
	line, err := ReadFile(filename)
	if err != nil {
		err = WriteFile(filename, "this is first line")
		if err != nil {
			fmt.Println("파일생성 에 실패했습니다", err)
			return
		}
		line, err = ReadFile(filename)
		if err != nil {
			fmt.Println("파일읽기 에 실패했습니다.", err)
			return
		}
	}
	fmt.Println("파일 내용 :", line)
}

type PasswordError struct {
	Len        int
	RequireLen int
}

func (err PasswordError) Error() string {
	return "암호길이 가 짧습니다."
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return PasswordError{len(password), 8}
	}
	return nil
}

// Error wrap

func multipleFromString(str string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(str))
	scanner.Split(bufio.ScanWords)

	pos := 0
	a, n, err := readNextInt(scanner)

	if err != nil {
		return 0, fmt.Errorf("Failed to readnextInt(), pos :%d,err:%w", pos, err)
	}

	pos += n + 1
	b, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos %d,err:%w", pos, err)
	}
	return a * b, nil
}

func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() {
		return 0, 0, fmt.Errorf("Failed to scan")
	}

	word := scanner.Text()
	number, err := strconv.Atoi(word)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to convert word to int, word %s, err :%w", word, err)
	}
	return number, len(word), nil
}

func readEq(eq string) {
	rst, err := multipleFromString(eq)
	if err != nil {
		fmt.Println(err)
		var numError *strconv.NumError
		if errors.As(err, &numError) {
			fmt.Println("number error", numError)
		}
	} else {
		fmt.Println(rst)
	}
}

func Ex02() {
	readEq("123 3")
	readEq("123 abc")
}

// panic 핸들링 하기 어려운 에러를 만났을떄 프래고르매 조기종료 하는 방법
// recover 는 지양 할것 문제 를 해결해야지 그냥 넘어가면 안됨

func divide(a, b int) {
	if b == 0 {
		panic("b can't declare as 0")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func f() {
	fmt.Println("f() 함수시작")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("복구 중 - ", r)
		}
	}()

	g()
	fmt.Println("f 함수 끝")
}

func g() {
	fmt.Printf("9/3 = %d \n", h(9, 3))
	fmt.Printf("9/3 = %d \n", h(9, 0))
}

func h(a, b int) int {
	if b == 0 {
		panic(" b should not be 0")
	}
	return a / b
}

func Ex03() {
	f()
	fmt.Printf("프로그램 실행중 ㅎㅎ")
}
