package chapter16

import (
	"fmt"
	"strings"
)

// Module and Packages
/**

 */

var (
	a = b + c
	b = f()
	c = f()
	d = 1
)

func Ex01() {
	abc := make(map[int]*strings.Builder)
	abc[0] = &strings.Builder{}
	abc[0].WriteRune('A')

	fmt.Println(abc[0].String())
}
func init() {
	d++
	fmt.Println("This is Init Function ", d)
}
func f() int {
	d++
	fmt.Println("f() d: ", d)
	return d
}
func Ex02() {
	fmt.Println("d : ", d)
}
