package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	encountered := map[int]int{}

	for index, num := range numbers {
		if otherIndex, ok := encountered[target-num]; ok {
			return []int{otherIndex + 1, index + 1}
		} else {
			encountered[num] = index
		}

	}

	return []int{}
}

func main() {

	numbers := []int{2, 7, 11, 15}
	target := 9

	fmt.Println("The first two numbers to add up to", target, "from the following numbers", numbers, "are", twoSum(numbers, target))

}
