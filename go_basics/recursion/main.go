package main

import "fmt"

// N-ый элемент арифметической прогрессии. Прямая рекурсия.
func ArithmeticProgressionElement(a1, d, n int) int {
	if n == 1 {
		return a1
	}
	return d + ArithmeticProgressionElement(a1, d, n-1)
}

// N-ый элемент геометрической прогрессии. Косвенная рекурсия.
func GeometricProgressionElement(b1, q, n int) int {
	if n == 1 {
		return b1
	}
	return q * ProgressionElement(b1, q, n)
}

func ProgressionElement(b1, q, n int) int {
	return GeometricProgressionElement(b1, q, n-1)
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println(ArithmeticProgressionElement(1, 2, 5))
	fmt.Println(GeometricProgressionElement(1, 2, 5))
	fmt.Println(fibonacci(42))
}
