package main

import (
	"fmt"
	"strconv"
	"strings"
)

func applyOperator(a float64, b float64, operator string) float64 {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	return 0
}

func evaluateExpression(expression string) (float64, error) {
	// replace colon with division operator
	expression = strings.ReplaceAll(expression, ":", "/")

	var tokens []string
	number := ""
	for _, char := range expression {
		if char >= '0' && char <= '9' || char == '.' {
			number += string(char)
		} else {
			if number != "" {
				tokens = append(tokens, number)
				number = ""
			}
			tokens = append(tokens, string(char))
		}
	}
	if number != "" {
		tokens = append(tokens, number)
	}

	// process multiplication and division
	var stack []string
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		if token == "*" || token == "/" {
			a, _ := strconv.ParseFloat(stack[len(stack)-1], 64)
			b, _ := strconv.ParseFloat(tokens[i+1], 64)
			// apply the operator
			result := applyOperator(a, b, token)
			stack[len(stack)-1] = fmt.Sprintf("%f", result)
			i++
		} else {
			stack = append(stack, token)
		}
	}

	// process addition
	result, _ := strconv.ParseFloat(stack[0], 64)
	for i := 1; i < len(stack); i += 2 {
		operator := stack[i]
		b, _ := strconv.ParseFloat(stack[i+1], 64)
		result = applyOperator(result, b, operator)
	}

	return result, nil
}

func main() {
	expression := "5+5+5*5:5"
	result, err := evaluateExpression(expression)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("The result of %s is: %f\n", expression, result)
}
