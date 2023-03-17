package main

import "fmt"

func media(number1 float64, number2 float64, number3 float64) float64 {
	var total float64 = ((number1 * 2.0) + (number2 * 3.0) + (number3 * 5.0)) / 10.0
	return total
}

func main() {

	var number1 float64
	var number2 float64
	var number3 float64

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&number1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&number2)

	fmt.Print("Digite o terceiro número: ")
	fmt.Scan(&number3)

	fmt.Print("A média ponderada entre eles é: ", media(number1, number2, number3))
}
