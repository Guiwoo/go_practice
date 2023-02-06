package chapter7

func Ex01(a, b int) int {
	return a + b

}

func Ex02(a, b, c int) int {
	return (a + b + c) / 3
}

func Ex03(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func Ex04(a, b int) (rst int, success bool) {
	if b == 0 {
		rst = 0
		success = false
		return
	}
	rst = a / b
	success = true
	return
}

func Ex05_fibo(a int) int {
	if a < 3 {
		return 1
	}
	return Ex05_fibo(a-1) + Ex05_fibo(a-2)
}
