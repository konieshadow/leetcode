package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func towSum1(nums []int, target int) []int {
	dict := make(map[int]int, len(nums))
	for i, num := range nums {
		dict[num] = i
	}

	for i, num := range nums {
		if j, ok := dict[target-num]; ok && j != i {
			return []int{i, j}
		}
	}

	return []int{}
}

func main() {
	nums := []int{3, 2, 4}
	target := 6

	fmt.Println(towSum1(nums, target))
}
