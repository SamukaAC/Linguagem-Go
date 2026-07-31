package main

import "fmt"

func main() {

	var num1 int

	fmt.Scanln(&num1)

	if (num1 % 2) == 0 {
		fmt.Printf("Numero %d eh par\n", num1)
		fmt.Println("Fim do Programa!")
	} else {
		fmt.Printf("Numero %d nao eh par\n", num1)
		fmt.Println("Fim do Programa!")
	}
}
