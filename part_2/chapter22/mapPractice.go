package chapter22

import "fmt"

/**
맵 키:벨류 자료형 구조
*/
// 순서 보장을 안해주는 맵 입니다.

func Ex01() {
	m := make(map[string]int)

	m["a"] = 169
	m["b"] = 170
	m["c"] = 190

	for i, v := range m {
		fmt.Println(i, v)
	}
}
