package main

import "fmt"

func add(a, b []int32, p int) []int32 {
	var c []int32
	var lenDiff = int32(len(a) - len(b))
	if lenDiff > 0 {
		for lenDiff != 0 {
			b = append(b, 0)
			lenDiff = int32(len(a) - len(b))
		}
	}
	for lenDiff != 0 {
		a = append(a, 0)
		lenDiff = int32(len(a) - len(b))
	}
	var extraDigit int32 = 0
	for i := 0; i < len(a); i++ {
		c = append(c, (a[i]+b[i]+extraDigit)%int32(p))
		extraDigit = (a[i] + b[i] + extraDigit) / int32(p)
	}
	if extraDigit != 0 {
		c = append(c, extraDigit)
	}
	return c
}

func main() {
	b := []int32{2, 2}
	a := []int32{2}
	fmt.Println(add(a, b, 3))
}
