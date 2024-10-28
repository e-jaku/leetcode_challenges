package main

import "fmt"

func removeElement(nums []int, val int) int {
	count := 0
	var result []int
	for index, numVal := range nums {
		if val == numVal {
			count++
			nums[index] = -1
		} else {
			result = append(result, numVal)
		}
	}

	copy(nums, result)

	return len(nums) - count
}

func main() {
	nums := []int{1, 2, 3, 3, 4, 2}
	val := 2
	remaining := removeElement(nums, val)
	fmt.Printf("Val %d was removed, remaining %d vals, remaining values:%v \n", val, remaining, nums[:remaining])
}
