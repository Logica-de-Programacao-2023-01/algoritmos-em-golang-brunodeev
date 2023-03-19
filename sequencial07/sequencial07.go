package main

import "fmt"

func main() {

	var number int

	fmt.Print("Digite um número: ")
	fmt.Scan(&number)

	fmt.Print("Seu antecessor é ", number-1, " e seu sucessor é ", number+1)
}
