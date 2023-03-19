package main

import "fmt"

func main() {

	var price float64

	fmt.Print("Qual o valor do produto? R$")
	fmt.Scan(&price)

	fmt.Print("\nO valor do produto com desconto de 10% é: R$", discount(price))
}

func discount(price float64) float64 {
	var calculate float64 = price - ((10.0 / 100.0) * price)
	return calculate
}
