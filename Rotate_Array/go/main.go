package main

import "fmt"

/*
// the inefficient way O(n^2)

func rotate(nums []int, k int) {
	for i := 0; i < k%len(nums); i++ {
		last := nums[len(nums)-1]
		next := nums[0]
		for j := 1; j < len(nums); j++ {
			tmp := nums[j]
			nums[j] = next
			next = tmp
		}
		nums[0] = last
	}
}
*/

func reverse(nums []int, firstP int, secondP int) {
	for firstP < secondP {
		tmp := nums[firstP]
		nums[firstP] = nums[secondP]
		nums[secondP] = tmp
		firstP++
		secondP--
	}
}

func rotate(nums []int, k int) {
	if len(nums) > 0 {
		n := k % len(nums)
		reverse(nums, 0, len(nums)-1) // full revers

		reverse(nums, 0, n-1)

		reverse(nums, n, len(nums)-1)
	}
}

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6}
	k := 8

	rotate(nums, k)

	fmt.Println(nums, k%len(nums))
}
