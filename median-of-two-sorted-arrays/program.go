package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	medianI := (len(nums1) + len(nums2)) / 2
	isOdd := (len(nums1)+len(nums2))%2 != 0
	var currentNum int
	var lastNum int

	for i1, i2 := 0, 0; ; {
		if i1 < len(nums1) && (i2 >= len(nums2) || nums1[i1] < nums2[i2]) {
			lastNum = currentNum
			currentNum = nums1[i1]
			i1++
		} else if i2 < len(nums2) {
			lastNum = currentNum
			currentNum = nums2[i2]
			i2++
		}

		if i1+i2-1 >= medianI {
			if isOdd {
				return float64(currentNum)
			} else {
				return float64(currentNum+lastNum) / 2
			}
		}
	}
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}

	//nums1 := []int{}
	//nums2 := []int{2, 4, 5}

	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
