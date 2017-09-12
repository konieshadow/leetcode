package main

import (
	"fmt"
)

func isMatch(str string, pattern string) bool {
	var s = 0
	var p = 0
	var match = 0
	var starIdx = -1

	for s < len(str) {
		if p < len(pattern) && (pattern[p] == '?' || str[s] == pattern[p]) {
			s++
			p++
		} else if p < len(pattern) && pattern[p] == '*' {
			match = s
			starIdx = p
			p++
		} else if starIdx != -1 {
			p = starIdx + 1
			match++
			s = match
		} else {
			return false
		}
	}

	for p < len(pattern) && pattern[p] == '*' {
		p++
	}

	return p == len(pattern)
}

func main() {
	// var result = isMatch("abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb",
	// 	 "**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb")
	// println(result)

	// var result2 = isMatch("aaaa", "***a")
	// fmt.Println(result2)

	var result3 = isMatch("aa", "*")
	fmt.Println(result3)
}
