package main

func removeDuplicates(nums []int) int {
	duplicateIndex := 0
	double := false
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[duplicateIndex] {
			nums[duplicateIndex+1] = nums[i]
			duplicateIndex++
			double = false
		} else {
			if !double {
				nums[duplicateIndex+1] = nums[i]
				duplicateIndex++
				double = true
			}
		}

	}

	return duplicateIndex + 1
}

func main() {

}
