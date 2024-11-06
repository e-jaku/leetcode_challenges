package main

import "fmt"

func maxProfit(prices []int) int {
	biggestDiff := 0
	minVal := 0
	for index, price := range prices {
		if index == 0 {
			minVal = price
			continue
		}

		diff := price - minVal

		if diff > biggestDiff {
			biggestDiff = diff
			continue
		}

		if price < minVal {
			minVal = price
		}
	}

	return biggestDiff
}

func main() {

	prices := []int{7, 6, 4, 3, 1}

	fmt.Println("The biggest achievable profit is:", maxProfit(prices))

}
