package main

import (
	"fmt"
	"math"
)

func myAtoi(str string) int {
	INT_MIN := uint(math.Pow(2, 31))
	INT_MAX := INT_MIN - 1

	sign := 1
	var r uint
	empty := true

	for _, c := range str {
		if empty {
			if c == ' ' {
				continue
			} else if c == '+' {
				sign = 1
				empty = false
			} else if c == '-' {
				sign = -1
				empty = false
			} else if r == 0 && c == '0' {
				empty = false
			} else if c >= '0' && c <= '9' {
				r = r*10 + uint(c-'0')
				empty = false
			} else {
				return 0
			}
		} else {
			if r == 0 && c == '0' {
				continue
			} else if c >= '0' && c <= '9' {
				r = r*10 + uint(c-'0')
				if r >= INT_MIN && sign < 0 {
					return -int(INT_MIN)
				} else if r > INT_MAX {
					return int(INT_MAX)
				}
			} else {
				break
			}
		}
	}

	return int(r) * sign
}

func main() {
	fmt.Println(myAtoi(" -42"))
	fmt.Println(myAtoi("-0042"))
	fmt.Println(myAtoi("42 wih words"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("9223372036854775808"))
	fmt.Println(myAtoi("-9223372036854775809"))
	fmt.Println(myAtoi("18446744073709551617"))
	fmt.Println(myAtoi(".1"))
	fmt.Println(myAtoi("-2147483648"))
}
