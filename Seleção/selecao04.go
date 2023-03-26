package main

import (
	"fmt"
	"math"
)

func main() {
	var weight float64
	var height float64
	var gender string

	fmt.Print("\nInforme seu peso (em kg): ")
	fmt.Scan(&weight)

	fmt.Print("\nInforme sua altura (em metros): ")
	fmt.Scan(&height)

	fmt.Print("\nInforme seu sexo (M para mulher e H para homem): ")
	fmt.Scan(&gender)

	if gender == "H" {
		if calculateImc(weight, height) < 20 {
			fmt.Println("\nVocê está abaixo do peso!")
		} else if calculateImc(weight, height) >= 20 && calculateImc(weight, height) <= 30 {
			fmt.Print("\nVocê está no peso ideal!")
		} else {
			fmt.Print("\nVocê está acima do peso!")
		}
	} else {
		if calculateImc(weight, height) < 19 {
			fmt.Print("\nVocê está abaixo do peso!")
		} else if calculateImc(weight, height) >= 19 && calculateImc(weight, height) <= 29 {
			fmt.Print("\nVocê está no peso ideal!")
		} else {
			fmt.Print("\nVocê está acima do peso!")
		}
	}

}

func calculateImc(weight float64, height float64) float64 {
	var imc float64 = 0
	imc = weight / (height * height)
	response := math.Round(imc*100) / 100
	return response
}
