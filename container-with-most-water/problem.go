package main

import (
	"fmt"
)

func maxArea(height []int) int {
	maxA := 0

	for i := 1; i < len(height); i++ {
		for j := 0; j < i; j++ {
			area := (i - j) * min(height[i], height[j])
			fmt.Println(area)
			maxA = max(maxA, area)
		}
	}

	return maxA
}

func min(num1, num2 int) int {
	if num2 < num1 {
		return num2
	}
	return num1
}

func max(num1, num2 int) int {
	if num2 > num1 {
		return num2
	}
	return num1
}

func main() {
	height := []int{1, 2, 4, 3}
	result := maxArea(height)
	fmt.Println(result)
}
