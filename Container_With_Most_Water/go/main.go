package main

import (
	"fmt"
)

func maxArea(height []int) int {
	leftIndex := 0
	maxArea := 0
	rightIndex := len(height) - 1
	for leftIndex < rightIndex {
		h := height[leftIndex]
		if h > height[rightIndex] {
			h = height[rightIndex]
		}

		area := h * (rightIndex - leftIndex)

		if area > maxArea {
			maxArea = area
		}

		if height[leftIndex] < height[rightIndex] {
			leftIndex++
		} else {
			rightIndex--
		}
	}

	return maxArea
}

func main() {

	nums := []int{1, 2, 4, 3}
	fmt.Println("The max area for this array", nums, "is", maxArea(nums))

}
