package chapter6

import (
	"fmt"
	"math"
	"math/big"
)

func equal(a, b float64) bool {
	return math.Nextafter(a, b) == b
}

func Ex01() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Println(equal(a+b, c))
}

func Ex02() {
	a, _ := new(big.Float).SetString("0.1")
	b, _ := new(big.Float).SetString("0.2")
	c, _ := new(big.Float).SetString("0.3")

	d := new(big.Float).Add(a, b)
	fmt.Println(a, b, c, d)
	fmt.Println(c.Cmp(d))
}
