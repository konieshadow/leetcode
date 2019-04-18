package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	max := 0
	n := len(s)
	dict := make(map[uint8]int)

	for i, j := 0, 0; i < n-max && j < n; j++ {
		if j1, exist := dict[s[j]]; exist {
			if j1+1 > i {
				i = j1 + 1
			}
		}

		if j-i+1 > max {
			max = j - i + 1
		}
		dict[s[j]] = j
	}

	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("abcdcebb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("au"))
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))
}
