package main

import "fmt"

func removeElement(nums []int, val int) int {
	removedIndex := 0
	for _, numVal := range nums {
		if numVal != val {
			nums[removedIndex] = numVal
			removedIndex++
		}
	}

	return removedIndex
}

func main() {
	nums := []int{1, 2, 3, 3, 4, 2}
	val := 2
	remaining := removeElement(nums, val)
	fmt.Printf("Val %d was removed, remaining %d vals, remaining values:%v \n", val, remaining, nums[:remaining])
}
