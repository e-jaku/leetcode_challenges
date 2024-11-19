package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	minLength := 0
	cArrCount := 0
	count := 0
	for index, val := range nums {
		if minLength == 1 {
			return 1 // min that can be reached
		}
		if minLength == 0 {
			count += val
			cArrCount++
			if count >= target {
				minLength = cArrCount
				cArrCount = 0
				count = 0
			}
		}

		if minLength != 0 && minLength > 1 {
			cArrCount = 0
			count = 0
			for i := index; i >= index-(minLength-2); i-- {
				count += nums[i]
				cArrCount++
				if count >= target {
					if cArrCount < minLength {
						minLength = cArrCount
					}
					break
				}
			}
		}

	}

	return minLength
}

func main() {
	target := 4
	nums := []int{1, 4, 4}
	fmt.Println("The min length of a subarray to sum up to", target, "with the following array", nums, "is", minSubArrayLen(target, nums))
}
