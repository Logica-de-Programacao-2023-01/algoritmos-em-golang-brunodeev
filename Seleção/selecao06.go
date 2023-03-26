package main

import (
	"fmt"
)

func main() {

	var number1 int
	var number2 int

	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&number1)

	fmt.Print("Digite outro número inteiro: ")
	fmt.Scan(&number2)

	if number1 > 0 && number2 > 0 {
		fmt.Println("A multiplicação desses números é: ", mult(number1, number2))
	} else {
		fmt.Println("A soma desses números é: ", sum(number1, number2))
	}

}

func mult(number1, number2 int) int {
	var result int = number1 * number2
	return result
}

func sum(number1, number2 int) int {
	var result int = number1 + number2
	return result
}
