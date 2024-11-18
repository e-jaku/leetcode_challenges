package main

import "fmt"

func romanToInt(s string) int {
	romanToInteger := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	sRunes := []rune(s)
	nextChar := -1
	count := 0
	for index, char := range sRunes {
		if index <= nextChar {
			continue
		}

		if index+1 <= len(sRunes)-1 {
			// check if 2 runes are mapped
			if ok := romanToInteger[fmt.Sprintf("%s%s", string(char), string(sRunes[index+1]))]; ok == 0 {
				// not a particular mapping we use only char
				val := romanToInteger[string(char)]
				count += val
			} else {
				nextChar = index + 1
				val := romanToInteger[fmt.Sprintf("%s%s", string(char), string(sRunes[index+1]))]
				count += val
			}
		} else {
			val := romanToInteger[string(char)]
			count += val
		}
	}
	return count
}

func main() {
	str := "MCMXCI"
	fmt.Println("The roman number", str, "converts to", romanToInt(str))
}
