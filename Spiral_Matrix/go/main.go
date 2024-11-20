package main

import "fmt"

func spiralOrder(matrix [][]int) []int {

	left := 0
	top := 0
	bottom := len(matrix)

	if bottom < 1 {
		return []int{}
	}

	right := len(matrix[0])
	res := []int{}

	for left < right && top < bottom {
		for i := left; i < right; i++ {
			res = append(res, matrix[top][i]) // add all moving from left to right
		}
		top++ // increase to avoid reiterating same row

		for i := top; i < bottom; i++ {
			res = append(res, matrix[i][right-1]) // add all moving from top to bottom
		}
		right-- // shrink to avoid reiterating same col

		if !(left < right && top < bottom) {
			return res
		}

		for i := right - 1; i >= left; i-- {
			res = append(res, matrix[bottom-1][i]) // add all moving from right to left at bottom row
		}

		bottom--

		for i := bottom - 1; i >= top; i-- {
			res = append(res, matrix[i][left]) // add all moving bottom to top from bottom left corner
		}
		left++
	}

	return res
}

func main() {

	matrix := [][]int{
		[]int{1, 2},
		[]int{3, 4},
	}

	fmt.Println("The spiral order of the matrix", matrix, "is", spiralOrder(matrix))

}
