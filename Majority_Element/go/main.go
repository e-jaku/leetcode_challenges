package main

import "fmt"

func majorityElement(nums []int) int {
	count := 0
	majorityElem := 0
	for _, val := range nums {
		if count == 0 {
			majorityElem = val
			count++
		} else {
			if val == majorityElem {
				count++
			} else {
				count--
			}
		}
	}
	return majorityElem
}

func main() {
	fmt.Println("Element occurring n/2 times is:", majorityElement([]int{1, 2, 2, 2, 4}))
}
