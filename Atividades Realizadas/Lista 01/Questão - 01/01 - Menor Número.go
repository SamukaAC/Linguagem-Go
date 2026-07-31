package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Scanln(&num1, &num2)

	if num1 > num2 {
		fmt.Printf("Numero %d, eh maior que %d\n", num1, num2)
	} else {
		fmt.Printf("Numero %d, eh maior que %d\n", num2, num1)
	}
}
