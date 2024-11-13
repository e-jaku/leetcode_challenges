package main

import "fmt"

func candy(ratings []int) int {
	if len(ratings) == 1 {
		return 1
	}

	countCandy := 0
	candyDistribution := make([]int, len(ratings))

	for index, val := range ratings {
		if index == 0 {
			next := ratings[index+1]

			if val > next {
				candyDistribution[index] = 2
				countCandy += 2
			} else {
				candyDistribution[index] = 1
				countCandy++
			}
			continue
		}

		if index == len(ratings)-1 {
			prev := ratings[index-1]
			if val > prev {
				prevCandy := candyDistribution[index-1]
				candyDistribution[index] = prevCandy + 1
				countCandy += prevCandy + 1
			} else {
				candyDistribution[index] = 1
				countCandy++
			}
			continue
		}

		prev := ratings[index-1]
		next := ratings[index+1]

		if val > next && val < prev && candyDistribution[index-1] <= 2 {
			//need backtracking
			candyDistribution[index] = 2
			countCandy += 2
			countCandy += backtrack(candyDistribution, ratings, index)
			continue
		}

		if val > prev {
			prevCandy := candyDistribution[index-1]
			candyDistribution[index] = prevCandy + 1
			countCandy += prevCandy + 1
			continue
		}

		if val <= next && val <= prev {
			candyDistribution[index] = 1
			countCandy += 1
			continue
		} else {
			candyDistribution[index] = 2
			countCandy += 2
			continue
		}
	}

	return countCandy
}

func backtrack(candyDistribution []int, ratings []int, index int) int {
	count := 0
	for i := index; i > 0; i-- {
		current := ratings[i]
		prevValue := ratings[i-1]

		if current < prevValue {
			// only increase previous value by one if smaller still

			if candyDistribution[i] >= candyDistribution[i-1] {
				candyDistribution[i-1] += 1
				count++
			}

		} else {
			break
		}
	}

	return count
}
func main() {

	ratings := []int{1, 2, 3, 5, 4, 3, 2, 1}

	fmt.Println("We need:", candy(ratings), "candy for children with the following ratings", ratings)
}
