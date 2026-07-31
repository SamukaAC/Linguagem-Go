package main

import "fmt"

func main() {

	var num1, num2 int

	fmt.Scanln(&num1, &num2)

	if (num1 + num2) == 42 {
		fmt.Printf("A soma de %d e %d eh equivalente ao enigma do universo\n", num1, num2)
	} else {
		fmt.Printf("A soma de %d e %d nao eh equivalente ao enigma do universo\n", num1, num2)
	}
	fmt.Println("System Down")
}
