package chapter15

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

/**
문자열
문자 집합 ASCII 문자 한개가 1바이트 0 ~ 255 8bits
CODE => 숫자에 의미 부여 Rune 이 그래서 그냥 숫자

UTF-8 문자코드 1~3 바이트 코드로 표현 (1~4바이트가 정의 임)
왜 가장 효율적인가 .. 적은거는 적게 희소성 있는 문자 는 바이트 를 높게 측정해서
UTF-16 은 기존 ASCII 랑 호환이 잘안되서 항상 변환이 필요하다
*/

func Ex01() {
	poet1 := "Holy moly"
	poet2 := `죽는 날까지 하늘을 우러러 블라블라`
	fmt.Println(poet2, poet1)
}

func Ex02() {
	str := "Hello world"

	//len => 바이트 길이를 반환 해줌
	for i := 0; i < len(str); i++ {
		fmt.Println(i)
		fmt.Printf("Type :%T,Value %d, Character : %c\n", str[i], str[i], str[i])
	}
}

func Ex03() {
	str := "Hello World"
	arr := []rune(str)
	arr[2] = rune(65)
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Type :%T,Value %d, Character : %c\n", arr[i], arr[i], arr[i])
	}
}

func Ex04() {
	str := "Hello World"

	for _, v := range str {
		fmt.Println(v)
	}
}

func Ex05() {
	str1 := "Hello World"
	str2 := str1

	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))

	fmt.Println(stringHeader1)
	fmt.Println(stringHeader2)
}

func Ex06() {
	var str string = "Hello world"
	var slice []byte = []byte(str)

	stringheader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	sliceHeader := (*reflect.StringHeader)(unsafe.Pointer(&slice))

	fmt.Printf("string : %x \n", stringheader.Data)
	fmt.Printf("slice : %x \n", sliceHeader.Data)
}

func toUpper(str string) string {
	var rst string
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			rst += string('A' + (c - 'a'))
		} else {
			rst += string(c)
		}
	}
	return rst
}

func toUpper2(str string) string {
	var builder strings.Builder
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			builder.WriteRune('A' + (c - 'a'))
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}
