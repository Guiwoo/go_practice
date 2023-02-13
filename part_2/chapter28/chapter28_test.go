package chapter28

import "testing"

func TestSquare(t *testing.T) {
	rst := square(9)
	if rst != 81 {
		t.Errorf("square(9) should be 81 but returns %d", rst)
	}
}

func TestSquare2(t *testing.T) {
	rst := square(3)

	if rst != 9 {
		t.Errorf("square(9) should be 81 but returns %d", rst)
	}
}

func BenchmarkFibo1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibo(20)
	}
}
func BenchmarkFibo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibo2(20)
	}
}
