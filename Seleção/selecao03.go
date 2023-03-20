package main

import "fmt"

func main() {
	var number int

	fmt.Print("Digite um número: ")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Println("Esse número é par!")
	} else {
		fmt.Println("Esse número é ímpar!")
	}
}
