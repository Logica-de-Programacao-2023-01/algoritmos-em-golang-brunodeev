package main

import (
	"fmt"
)

func main() {

	var number int

	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&number)

	if number%3 == 0 && number%5 == 0 {
		fmt.Println("O número digitado é múltiplo de 5 e 3 ao mesmo tempo!")
	} else {
		fmt.Println("O número digitado não é múltiplo de 5 e 3!")
	}

}
