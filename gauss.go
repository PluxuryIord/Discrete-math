package main

import "fmt"

type frac struct {
	a, b int64
}

func gcd(a int64, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int64) int64 {
	if a == 0 || b == 0 {
		return a + b
	}
	return (a * b) / gcd(a, b)
}

func fracAdd(x, y frac, operation string) frac {
	result := frac{0, 1}
	switch operation {
	case "+":
		result.a = x.a*y.b + x.b*y.a
		result.b = x.b * y.b
	case "-":
		result.a = x.a*y.b - x.b*y.a
		result.b = x.b * y.b
	case "*":
		result.a = x.a * y.a
		result.b = x.b * y.b
	case "/":
		result.a = x.a * y.b
		result.b = x.b * y.a
	}
	return result
}

func gauss(matrix [][]frac, b []frac) []frac {
	n := len(matrix)
	for i := 0; i < n; i++ {
		pivot := matrix[i][i]
		for j := i + 1; j < n; j++ {
			ratio := fracAdd(matrix[j][i], pivot, "/")
			for k := i; k < n; k++ {
				matrix[j][k] = fracAdd(matrix[j][k], fracAdd(ratio, matrix[i][k], "*"), "-")
			}
			b[j] = fracAdd(b[j], fracAdd(ratio, b[i], "*"), "-")
		}
	}

	x := make([]frac, n)
	temp := frac{0, 1}
	for i := n - 1; i >= 0; i-- {
		temp = b[i]
		for j := i + 1; j < n; j++ {
			temp = fracAdd(temp, fracAdd(matrix[i][j], x[j], "*"), "-")
		}
		x[i] = fracAdd(temp, matrix[i][i], "/")
	}
	return x
}

func main() {
	matrix := [][]frac{
		{{-4, 1}, {-1, 1}, {8, 1}},
		{{7, 1}, {-7, 1}, {7, 1}},
		{{5, 1}, {-1, 1}, {-4, 1}},
	}
	b := []frac{{2, 1}, {3, 1}, {7, 1}}
	solution := gauss(matrix, b)
	fmt.Printf("Решение: %v\n", solution)
}
