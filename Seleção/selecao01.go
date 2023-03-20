package main

import "fmt"

func main() {
	var number1 int
	var number2 int

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&number1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&number2)

	if number1 > number2 {
		fmt.Println("O número ", number1, " é maior que o número ", number2)
	} else if number2 > number1 {
		fmt.Println("O número ", number2, " é maior que o número ", number1)
	} else {
		fmt.Println("Ih, mané! Os dois são iguais :(")
	}
}
