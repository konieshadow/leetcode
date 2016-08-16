package main

func reverseVowels(s string) string {
	arrVowelChar := []byte{}
	arrVowelPos := []int{}

	for i := 0; i < len(s); i++ {
		if (isVowel(s[i])) {
			arrVowelChar = append(arrVowelChar, s[i])
			arrVowelPos = append(arrVowelPos, i)
		}
	}

	arrS := []byte(s)

	for i := 0; i < len(arrVowelPos); i++ {
		arrS[arrVowelPos[i]] = arrVowelChar[len(arrVowelPos) -i -1]
	}

	return string(arrS)
}

func isVowel(c byte) bool {
	if (c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
		c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U') {
		return true
	}

	return false
}

func main() {
	s1 := "hello"
	println(reverseVowels(s1))

	s2 := "leetcode"
	println(reverseVowels(s2))
}