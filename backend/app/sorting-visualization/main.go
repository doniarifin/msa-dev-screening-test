package main

import (
	"fmt"
)

func main() {
	// Example input
	data := []int{1, 4, 5, 6, 8, 2}

	printVerticalBarchart(data)
}

func printVerticalBarchart(data []int) {
	maxHeight := max(data)

	// Generate the barchart
	for level := maxHeight; level > 0; level-- {
		for _, value := range data {
			if value >= level {
				fmt.Print("| ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}

	// Print the horizontal axis with data values
	for _, value := range data {
		fmt.Printf("%d ", value)
	}
	fmt.Println()
}

func max(data []int) int {
	maxValue := data[0]
	for _, value := range data {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}
