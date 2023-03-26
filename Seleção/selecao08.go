package main

import (
	"fmt"
)

func main() {

	var day float64

	fmt.Print("\nDigite um número de 1 a 7 para representar um dia na semana: ")
	fmt.Scan(&day)

	switch day {
	case 1:
		fmt.Print("\nDomingo!\n")
	case 2:
		fmt.Print("\nSegunda-Feira!\n")
	case 3:
		fmt.Print("\nTerça-Feira!\n")
	case 4:
		fmt.Print("\nQuarta-Feira!\n")
	case 5:
		fmt.Print("\nQuinta-Feira!\n")
	case 6:
		fmt.Print("\nSexta-Feira!\n")
	case 7:
		fmt.Print("\nSábado!\n")
	default:
		fmt.Print("\nO número digitado é inválido!\n")
	}
}
