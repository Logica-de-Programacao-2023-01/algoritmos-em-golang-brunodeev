package main

import (
	"fmt"
	"math"
)

func main() {

	var weight float64
	var height float64

	fmt.Print("Informe seu peso (em kg): ")
	fmt.Scan(&weight)

	fmt.Print("Informe sua altura (em metros): ")
	fmt.Scan(&height)

	fmt.Println("\nSeu IMC é ", calculateImc(weight, height))
}

func calculateImc(weight float64, height float64) float64 {
	var imc float64 = 0
	imc = weight / (height * height)
	response := math.Round(imc*100) / 100
	return response
}
