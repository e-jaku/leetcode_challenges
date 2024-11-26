package main

import "fmt"

func longestSubarray(nums []int) int {
	max := 0
	zerosCanceled := 0
	startIndex := 0
	for index, val := range nums {
		if val != 1 {
			// found a 0 increment the counter, we ignore at most one zero otherwise we adapt the startIndex
			zerosCanceled++
		}

		for zerosCanceled > 1 { // we found a second zero now we need to adapt the start index
			if nums[startIndex] == 0 { // we increase the start index until we find the previous zero we canceled
				zerosCanceled--
			}
			startIndex++
		}

		if index-startIndex >= max { // if we found a new longest sequence we save it
			max = index - startIndex
		}
	}

	return max
}

func main() {
	nums := []int{0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 1, 1}
	fmt.Println("The longest subarray of array", nums, "is of length", longestSubarray(nums))
}
