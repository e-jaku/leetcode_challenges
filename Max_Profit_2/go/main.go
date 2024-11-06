package main

import "fmt"

func maxProfit(prices []int) int {
	totalDiff := 0
	minVal := 0
	for index, price := range prices {
		if index == 0 {
			minVal = price
			continue
		}

		diff := price - minVal

		if diff > 0 {
			totalDiff += diff
			minVal = price // we assume re-buying directly after
			continue
		}

		if price < minVal {
			minVal = price
		}
	}

	return totalDiff
}

func main() {

	prices := []int{7, 6, 4, 3, 1}

	fmt.Println("The biggest achievable profit is:", maxProfit(prices))

}
