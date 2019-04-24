package main

import (
	"fmt"
)

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	matrix := make([][]int32, 0)

	result := make([]int32, len(s))

	row := 0
	col := 0
	mode := 1

	for _, c := range s {
		if len(matrix) < col+1 {
			matrix = append(matrix, make([]int32, numRows))
		}
		matrix[col][row] = c

		switch mode {
		case 1:
			if row == numRows-1 {
				mode = 2
				row--
				col++
			} else {
				row++
			}
		case 2:
			if row == 0 {
				mode = 1
				row++
			} else {
				row--
				col++
			}
		}
	}

	index := 0
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(matrix); j++ {
			if matrix[j][i] > 0 {
				result[index] = matrix[j][i]
				index++
			}
		}
	}

	return string(result)
}

func main() {
	//fmt.Println(convert("LEETCODEISHIRING", 3))
	//fmt.Println(convert("LEETCODEISHIRING", 4))
	fmt.Println(convert("AB", 1))
}
