package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int)  {
    i := m - 1
    j := n - 1
    k := m + n - 1

    for {
        println(i, j , k)

        if i < 0 {
            nums1[k] =nums2[j]
            j = j - 1
        } else if j < 0 {
            nums1[k] = nums1[i]
            i = i - 1
        } else {
            if nums1[i] > nums2[j] {
                nums1[k] = nums1[i]
                i = i - 1
            } else {
                nums1[k] =nums2[j]
                j = j - 1
            }
        }
        
        k = k - 1
        if (k < 0) {
            break
        }
    }
}

func main() {
    // nums1 := []int{1, 2, 3, 0, 0, 0, 0, 0, 0, 0}
    // nums2 := []int{2, 5, 7, 9}
    // merge(nums1, 3, nums2, 4)
    
    nums1 := []int{0}
    nums2 := []int{1}
    merge(nums1, 0, nums2, 1)

    fmt.Printf("%v", nums1)
}