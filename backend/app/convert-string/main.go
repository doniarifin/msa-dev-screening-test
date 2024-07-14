package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	str := "1AB23C5678D"
	arrayStr := stringToArray(str)
	arrayInt := stringToIntArray(str)
	arrayStr2 := stringToCharArray(str)

	fmt.Println(arrayStr)
	fmt.Println(arrayInt)
	fmt.Println(arrayStr2)
}

func stringToArray(input string) []string {
	result := strings.Split(input, "")
	return result
}

func stringToIntArray(input string) []int {
	var result []int
	for _, char := range input {
		if unicode.IsDigit(char) {
			num, _ := strconv.Atoi(string(char))
			result = append(result, num)
		}
	}
	return result
}

func stringToCharArray(input string) []string {
	var result []string
	for _, char := range input {
		if unicode.IsLetter(char) {
			result = append(result, string(char))
		}
	}
	return result
}
