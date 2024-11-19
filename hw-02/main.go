package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var op string

	fmt.Print("")
	fmt.Scanln(&num1)
	fmt.Print("")
	fmt.Scanln(&op)
	fmt.Print("")
	fmt.Scanln(&num2)

	var result float64
	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		fmt.Println("Не тот знак")
		return
	}

	fmt.Println("Ответ", result)
}
