package main

import (
	"fmt"
	"time"
)

func visualizeArray(arr []int) {
	maxVal := findMax(arr)
	for i := maxVal; i > 0; i-- {
		for _, v := range arr {
			if v >= i {
				fmt.Print("| ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println("\n")
}

// Helper function to find the maximum value in the array
func findMax(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key

		fmt.Println("Step", i)
		visualizeArray(arr)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	arr := []int{1, 4, 5, 6, 8, 2}
	fmt.Println("Initial array:")
	visualizeArray(arr)
	time.Sleep(2 * time.Second)

	fmt.Println("Sorting steps:")
	insertionSort(arr)

	fmt.Println("Sorted array:")
	visualizeArray(arr)
}
