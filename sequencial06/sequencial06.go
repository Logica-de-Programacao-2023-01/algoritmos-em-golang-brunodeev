package main

import "fmt"

func main() {

	var salary float64

	fmt.Print("Digite seu salário: ")
	fmt.Scan(&salary)

	fmt.Print("\nSeu salário com aumento de 15% é: R$", newSalary(salary))
}

func newSalary(salary float64) float64 {
	var calculate float64 = salary + ((15.0 / 100.0) * salary)
	return calculate
}
