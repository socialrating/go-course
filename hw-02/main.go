package main

import (
	"fmt"

	"github.com/socialrating/go-course/hw-02/calc"
)

func main() {
	var num1, num2 float64
	var operatorstr string

	fmt.Print("первое число: ")
	fmt.Scanln(&num1)

	fmt.Print("строчный оператор: ")
	fmt.Scanln(&operatorstr)

	fmt.Print("второе число: ")
	fmt.Scanln(&num2)

	calc := calc.NewCalculator()
	result := calc.Calculate(num1, num2, operatorstr)
	fmt.Println("Ответ:", result)
}
