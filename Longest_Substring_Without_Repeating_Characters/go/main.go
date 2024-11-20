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
	for index, val := range s {
		if kIndex, ok := encountered[val]; !ok {
			//new char
			encountered[val] = index
		} else {
			if counter > maxLen {
				maxLen = counter
			}
			counter = (index - kIndex - 1)
			deleteBeforeVal(kIndex, encountered)
			encountered[val] = index
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

func deleteBeforeVal(index int, known map[rune]int) {
	for key, val := range known {
		if val < index {
			delete(known, key)
		}
	}
}

func main() {
	str := "abcabcbb"
	fmt.Println("The longest substring without repeating characters of string", str, "is", lengthOfLongestSubstring(str))
}
