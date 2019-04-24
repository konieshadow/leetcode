package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	r := 0
	sign := 1
	if x < 0 {
		sign = -1
	}
	x1 := x / sign

	for x1 > 0 {
		r = r*10 + x1%10
		x1 = x1 / 10
	}

	r = r * sign
	if r < -int(math.Pow(2, 31)) || r > int(math.Pow(2, 31))-1 {
		return 0
	}

	return r
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(-123))
}
