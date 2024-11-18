package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}

	if len(s) == len(t) {
		return s == t
	}

	if len(s) == 0 {
		return true
	}

	sIndex := 0
	sRunes := []rune(s)
	for _, char := range t {
		if sIndex <= len(sRunes)-1 {
			if char == sRunes[sIndex] {
				sIndex++
			}
			if sIndex > len(sRunes)-1 {
				return true
			}
		} else {
			break
		}
	}

	return false
}

func main() {
	s := ""
	t := "ahbgdc"

	fmt.Println("The string", s, "is a subsequence of the string", t, isSubsequence(s, t))
}
