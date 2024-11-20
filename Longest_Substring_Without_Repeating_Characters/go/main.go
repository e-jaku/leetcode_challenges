package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	encountered := map[rune]int{}
	maxLen := 0
	counter := 0
	ignoreIndex := 0
	for index, val := range s {
		if kIndex, ok := encountered[val]; !ok {
			//new char
			encountered[val] = index
		} else {
			encountered[val] = index
			if kIndex < ignoreIndex {
				counter++
				continue
			}

			if counter > maxLen {
				maxLen = counter
			}
			counter = (index - kIndex - 1)
			ignoreIndex = kIndex
			if index == len(s)-1 {
				break
			}
		}

		counter++
	}

	if counter > 0 && counter > maxLen {
		return counter
	}

	return maxLen
}

func main() {
	str := "aabaab!bb"
	fmt.Println("The longest substring without repeating characters of string", str, "is", lengthOfLongestSubstring(str))
}
