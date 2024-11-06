package main

import "fmt"

func canJump(nums []int) bool {
	n := len(nums) - 1 // end of slice
	currentPossible := 0
	for index, num := range nums {
		if index == 0 {
			currentPossible = num
		}

		if currentPossible > 0 || num > 0 {
			if currentPossible <= num {
				currentPossible = num
			}
		}

		if currentPossible == 0 && index != n {
			return false
		}
		currentPossible--
	}

	return true
}

func main() {
	nums := []int{3, 2, 1, 0, 4}
	fmt.Println("Based on slice:", nums, "we reach the end:", canJump(nums))
}
