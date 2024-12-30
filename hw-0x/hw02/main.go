package main

import (
	"fmt"
)

func main() {
	var sum int
	for i := 0; i <= 50; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println("Сумма чётных чисел от 0 до 50:", sum)
}
