package main

import "fmt"

func main() {

	var number int

	fmt.Print("Digite um número: ")
	fmt.Scan(&number)

	for i := 1; i <= number; i++ {
		if number%i == 0 {
			fmt.Println(i)
		}
	}
}
