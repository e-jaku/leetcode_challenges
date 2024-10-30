package main

import "fmt"

func removeDuplicates(nums []int) int {
	uniqueIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[uniqueIndex] {
			nums[uniqueIndex+1] = nums[i]
			uniqueIndex++
		}
	}

	return uniqueIndex + 1
}

func main() {
	fmt.Println("Found", removeDuplicates([]int{1, 2, 3, 3, 4}), "unique values after removing duplicates")
}
