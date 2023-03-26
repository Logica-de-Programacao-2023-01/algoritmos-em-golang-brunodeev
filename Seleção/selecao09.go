package main

import (
	"fmt"
)

func main() {

	var number1 float64
	var number2 float64
	var number3 float64

	fmt.Print("\nDigite um número: ")
	fmt.Scan(&number1)

	fmt.Print("\nDigite outro número: ")
	fmt.Scan(&number2)

	fmt.Print("\nDigite mais um número: ")
	fmt.Scan(&number3)

	if number1 < number2 && number1 < number3 {
		if number2 < number3 {
			fmt.Println(number1)
			fmt.Println(number2)
			fmt.Println(number3)
		} else {
			fmt.Println(number1)
			fmt.Println(number3)
			fmt.Println(number2)
		}
	} else if number2 < number1 && number2 < number3 {
		if number1 < number3 {
			fmt.Println(number2)
			fmt.Println(number1)
			fmt.Println(number3)
		} else {
			fmt.Println(number2)
			fmt.Println(number3)
			fmt.Println(number1)
		}
	} else if number3 < number1 && number3 < number2 {
		if number1 < number2 {
			fmt.Println(number3)
			fmt.Println(number1)
			fmt.Println(number2)
		} else {
			fmt.Println(number3)
			fmt.Println(number2)
			fmt.Println(number1)
		}
	} else {
		fmt.Println("Números se repetem!")
	}
}
