package main

import (
	"fmt"
)

func main() {

	var age int

	fmt.Print("\nDigite sua idade: ")
	fmt.Scan(&age)

	if age <= 9 {
		fmt.Print("\nMirim")
	} else if age <= 13 {
		fmt.Print("\nInfantil")
	} else if age <= 17 {
		fmt.Print("\nJuvenil")
	} else {
		fmt.Print("\nAdulto")
	}
}
