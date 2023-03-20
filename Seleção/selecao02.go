package main

import "fmt"

func main() {
	var number1 int
	var number2 int
	var number3 int

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&number1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&number2)

	fmt.Print("Digite o terceiro número: ")
	fmt.Scan(&number3)

	if number1 < number2 && number1 < number3 {
		fmt.Print("O número ", number1, " é o menor deles!")
	} else if number2 < number1 && number2 < number3 {
		fmt.Print("O número ", number1, " é o menor deles!")
	} else if number3 < number1 && number3 < number2 {
		fmt.Print("O número ", number1, " é o menor deles!")
	} else if number1 == number2 && number1 != number3 {
		fmt.Print("Há números que se repetem!")
	} else if number1 == number3 && number1 != number2 {
		fmt.Print("Há números que se repetem!")
	} else if number2 == number3 && number2 != number1 {
		fmt.Print("Há números que se repetem!")
	} else if number1 == number2 && number2 == number3 {
		fmt.Print("Ih, mané! Os números são iguais")
	} else {
		fmt.Print("Error 404")
	}
}
