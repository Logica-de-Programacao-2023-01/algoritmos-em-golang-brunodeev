package main

import "fmt"

func main() {
	var number int

	fmt.Print("Digite um número: ")
	fmt.Scan(&number)

	for i := 1; i <= 10; i++ {
		fmt.Println(i, "x", number, ": ", number*i)
	}
}
