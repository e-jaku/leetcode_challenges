package main

import (
	"fmt"
)

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	// water trapped is determined from the highest bar to the left and the highest bar to the right of a water pool
	// so for example if we have [0 1 0 2 0] the highest bar to the left when coming from the left is 1
	// and the highest bar to the right when coming from the right is 2

	// we iterate array three times, once from left to right and fill left heights, once from the right and fill right heights
	// we then calculate the difference between left and right at a specific position

	leftViewedHeights := make([]int, len(height))
	rightViewedHeights := make([]int, len(height))

	maxHeight := 0
	for index := range leftViewedHeights {
		if height[index] > maxHeight {
			maxHeight = height[index]
		}
		leftViewedHeights[index] = maxHeight
	}

	maxHeight = 0
	for i := len(height) - 1; i >= 0; i-- {
		if height[i] > maxHeight {
			maxHeight = height[i]
		}
		rightViewedHeights[i] = maxHeight
	}

	count := 0
	for index, val := range height {
		minHeight := min(leftViewedHeights[index], rightViewedHeights[index]) // min of the two since otherwise water would overflow
		diff := minHeight - val

		count += diff
	}

	return count
}

func main() {

	heights := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	fmt.Println("Based on these heights", heights, "we can capture", trap(heights), "water")
}
