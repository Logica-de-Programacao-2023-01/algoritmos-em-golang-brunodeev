package main

import "fmt"

func main() {
	var number int

	fmt.Print("Digite um número: ")
	fmt.Scan(&number)

	fmt.Println("O dobro:", double(number))
	fmt.Println("O triplo:", triple(number))
	fmt.Println("O quádruplo:", quadruple(number))

}

func double(number int) int {
	total := 0
	total = number * 2
	return total
}

func triple(number int) int {
	total := 0
	total = number * 3
	return total
}

func quadruple(number int) int {
	total := 0
	total = number * 4
	return total
}
