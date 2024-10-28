package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	result := append(nums1[:m], nums2...)
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	copy(nums1, result)
}

func main() {

	nums1 := []int{1, 6, 7, 0, 0, 0}
	nums2 := []int{2, 6, 8}
	m, n := 3, 3
	merge(nums1, m, nums2, n)

	fmt.Println(nums1)
}
