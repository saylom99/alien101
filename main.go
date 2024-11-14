package main

import "fmt"

func getValue(c rune) int {
	switch c {
	case 'A':
		return 1
	case 'B':
		return 5
	case 'Z':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'R':
		return 1000
	default:
		return 0
	}
}

func convertAndSumAlienNum(alienNum string) int {
	result := 0
	prevValue := 0
	runes := []rune(alienNum)

	for i := len(runes) - 1; i >= 0; i-- {
		currentValue := getValue(runes[i])

		if prevValue > currentValue {
			result = result - currentValue
		} else {
			result = result + currentValue
		}

		prevValue = currentValue
	}

	return result
}

func main() {
	tests := []string{"AAA", "LBAAA", "RCRZCAB"}
	for _, test := range tests {
		fmt.Printf("Input: %s\n", test)
		fmt.Printf("Output: %d\n", convertAndSumAlienNum(test))
	}
}
