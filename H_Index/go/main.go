package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	//sort citations
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	maxIndex := 0
	for index, val := range citations {
		if index < val {
			maxIndex = index
		} else {
			break
		}

	}

	if maxIndex >= len(citations) {
		return len(citations)
	}
	if citations[maxIndex] >= (maxIndex + 1) {
		return maxIndex + 1
	} else {
		if citations[maxIndex] > len(citations) {
			return len(citations)
		}
		return citations[maxIndex]
	}
}

func main() {
	citations := []int{1, 7, 9, 2}
	fmt.Println("An author with the following citations", citations, "has an H-Index of:", hIndex(citations))
}
