package main

import "fmt"

func checkPerfectNumber(num int) bool {
	if (num == 1) {
		return false
	}

	sum := 1
	divisor := 1
	
	for i := 2; i < num / i && i < num / divisor; i++ {
		fmt.Printf("i=%d\n", i)
		
		if (num % i != 0) {
			continue
		}

		divisor = i

		fmt.Printf("divisor=%d\n", divisor)
		fmt.Printf("divisor2=%d\n", num / divisor)
		
		sum = sum + divisor + num / divisor

		fmt.Printf("sum=%d\n", sum)
		if (sum > num) {
			return false
		}
	}

	return sum == num
}

func main() {
	var result = checkPerfectNumber(28)
	println(result)
}