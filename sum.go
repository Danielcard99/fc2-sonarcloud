package main

import "math"

func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func div(a int, b int) int {
	if b == 0 {
		return 0 // prevenção de divisão por zero
	}
	return a / b
}

func mod(a int, b int) int {
	if b == 0 {
		return 0
	}
	return a % b
}

func pow(a int, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sqrt(a int) int {
	return int(math.Sqrt(float64(a)))
}

func cube(a int) int {
	return a * a * a
}

func square(a int) int {
	return a * a
}

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Div(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

func IsEven(n int) bool {
	return n%2 == 0
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
