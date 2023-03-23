package main

import "fmt"

func main() {

	var number, sum, count int

	fmt.Println("Digite uma sequência de números inteiros (digite 0 para sair):")

	for {
		fmt.Scan(&number)

		if number == 0 {
			break
		}

		sum += number
		count++
	}

	if count > 0 {
		media := float64(sum) / float64(count)
		fmt.Printf("A média aritmética dos %d números é %.2f\n", count, media)
	} else {
		fmt.Println("Nenhum número foi digitado.")
	}
}
