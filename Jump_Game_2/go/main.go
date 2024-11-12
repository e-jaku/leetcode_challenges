package main

import (
	"fmt"
)

func jump(nums []int) int {
	count := 0
	nextIndex := 0
	for index, val := range nums {
		if index == len(nums)-1 {
			return count
		}

		if val+index >= len(nums)-1 {
			return count + 1
		}

		if index < nextIndex {
			continue
		}

		count++

		max := 0
		for i := 1; i <= val; i++ {
			if index+i >= len(nums)-1 {
				return count
			}
			num := nums[index+i]
			if max <= num+index+i {
				max = num + index + i
				nextIndex = index + i
			}

			if num+index+i >= len(nums)-1 {
				return count + 1
			}
		}
	}

	return count
}

func main() {
	nums := []int{3, 4, 3, 2, 5, 4, 3}
	fmt.Println("Based on slice:", nums, "we reach the end in:", jump(nums), "jumps")
}
