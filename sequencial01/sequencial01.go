package main

import "fmt"

func main() {

	var number1 int
	var number2 int
	var number3 int

	fmt.Print("Digite um número: ")
	fmt.Scan(&number1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&number2)

	fmt.Print("Digite o terceiro número: ")
	fmt.Scan(&number3)

	fmt.Println("A soma dos três valores digitados é ", number1+number2+number3)
}
