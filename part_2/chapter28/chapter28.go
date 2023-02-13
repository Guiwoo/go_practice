package chapter28

func square(i int) int {
	return i * i
}

func fibo(i int) int {
	if i < 1 {
		return 0
	}
	if i < 2 {
		return i
	}
	return fibo(i-2) + fibo(i-1)
}

func fibo2(i int) int {
	if i < 1 {
		return i
	}
	one := 1
	two := 0
	rst := 0
	for j := 2; j <= i; j++ {
		rst = one + two
		two = one
		one = rst
	}
	return rst
}
