package main

import (
	"fmt"
)

func main() {

	var salary float64

	fmt.Print("\nDigite seu salário atual: R$")
	fmt.Scan(&salary)

	if salary < 1000.0 {
		fmt.Printf("\nSeu novo salário será de: R$%.2f", tenPorcent(salary))
	} else {
		fmt.Printf("\nSeu novo salário será de: R$%.2f", fivePorcent(salary))
	}

}

func tenPorcent(salary float64) float64 {
	var result float64 = salary + (salary * 0.1)
	return result
}

func fivePorcent(salary float64) float64 {
	var result float64 = salary + (salary * 0.05)
	return result
}
