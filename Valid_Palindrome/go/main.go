package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	stripped := strip(s)
	lowerSt := strings.ToLower(stripped)
	runeArray := []rune(lowerSt)

	endIndex := len(runeArray) - 1
	for index, val := range runeArray {
		endVal := runeArray[endIndex]

		if val != endVal {
			return false
		}

		if endIndex <= index {
			//reached the middle
			break
		}

		endIndex--
	}

	return true

}

func strip(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') {
			result.WriteByte(b)
		}
	}
	return result.String()
}

func main() {
	str := "A man, a plan, a canal: Panama"
	fmt.Println("The string", str, "is a palindrome:", isPalindrome("A man, a plan, a canal: Panama"))
}
