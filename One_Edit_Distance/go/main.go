package main

import (
	"fmt"
	"math"
)

func isOneEditDistance(s, t string) bool {
	if math.Abs(float64(len(t)-len(s))) > 1 {
		return false // need to be one char apart
	}

	if s == t {
		return true
	}

	// reorder the strings so that first is the shorter one
	var first, second string

	if len(s) < len(t) {
		first = s
		second = t
	} else {
		first = t
		second = s
	}

	for index, char := range []byte(first) {
		if char != second[index] {
			// they differ, we have two options
			if len(first) == len(second) {
				return first[index+1:] == second[index+1:] // if the remainder is equal
			} else {
				return first[index:] == second[index+1:] //second is the longer one we skip the next char of second and compare the remainders
			}
		}
	}

	return len(second) == len(first)+1
}

func main() {
	str1 := "asf"
	str2 := "asfa"

	fmt.Println("The string", str1, "is one edit away from", str2, ":", isOneEditDistance(str1, str2))
}
