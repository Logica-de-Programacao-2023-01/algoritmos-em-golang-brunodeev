package main

import "fmt"

func ageDays(age int) int {
	var result int = age * 365
	return result
}

func main() {

	var age int

	fmt.Print("Digite sua idade: ")
	fmt.Scan(&age)

	fmt.Print("Sua idade em dias é: ", ageDays(age))
}
