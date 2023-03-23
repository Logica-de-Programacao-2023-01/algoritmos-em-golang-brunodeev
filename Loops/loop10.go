package main

import "fmt"

func main() {

	var number, big, count int

	fmt.Println("Digite uma sequência de números inteiros (digite 0 para sair):")

	for {
		fmt.Scan(&number)

		if number == 0 {
			break
		}

		if number >= big {
			big = number
		}

		count++
	}

	if count > 0 {
		fmt.Print("O maior dos ", count, " números é ", big, "\n")
	} else {
		fmt.Println("Nenhum número foi digitado.")
	}
}
