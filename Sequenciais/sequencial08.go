package main

import "fmt"

func main() {
	var days int
	var diary float64

	fmt.Print("Quantos dias foram trabalhados esse mês? ")
	fmt.Scan(&days)

	fmt.Print("\nQuanto é cobrado por dia? ")
	fmt.Scan(&diary)

	fmt.Print("Pode-se dizer que seu salário é de R$", salary(days, diary))
}

func salary(days int, diary float64) float64 {
	var result float64 = float64(days) * diary
	return result
}
