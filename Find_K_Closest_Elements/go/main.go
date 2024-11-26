package main

import "fmt"

func findClosestElements(arr []int, k int, x int) []int {
	var start, end int
	end = len(arr) - k // start index can go at most up to len(arr) -k otherwise it will not contain k elements

	for start < end {
		m := (start + end) / 2 // calculate middle of the sub array, we will find a middle which is the closest to the value x

		if x-arr[m] <= arr[m+k]-x {
			end = m // closer to the lower end
		} else {
			start = m + 1 // closer to the higher end
		}
	}

	return arr[start : start+k] // return subarray starting form start and spanning k elements (array is already sorted)
}

func main() {

	k := 2
	x := 2
	arr := []int{-1, 0, 1, 2, 5}

	fmt.Println("The", k, "nearest elements to", x, "from", arr, "are", findClosestElements(arr, k, x))
}
