package main

import "fmt"

func main() {

	var weight float64

	fmt.Print("Qual o seu peso (em kg)? ")
	fmt.Scan(&weight)

	fmt.Print("\nSeu peso em libras é: ", convert(weight), " lbs")
}

func convert(weight float64) float64 {
	var result float64 = weight * 2.20462
	return result
}
