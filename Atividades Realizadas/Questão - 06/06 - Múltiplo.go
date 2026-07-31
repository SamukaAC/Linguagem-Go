package main

import "fmt"

func main() {

	var num1, num2, multi int

	fmt.Scanln(&num1, &num2)

	if num1 > num2 {
		multi = (num1 % num2)
		if multi == 0 {
			fmt.Printf("Numero: %d\nÉ múltiplo de: %d\n", num1, num2)
		} else {
			fmt.Printf("Numero: %d\nNão é múltiplo de: %d\n", num1, num2)
		}
	} else if num2 > num1 {
		multi = (num2 % num1)
		if multi == 0 {
			fmt.Printf("Numero: %d\nÉ múltiplo de: %d\n", num2, num1)
		} else {
			fmt.Printf("Numero: %d\nNão é múltiplo de: %d\n", num1, num2)
		}
	} else {
		fmt.Printf("Numero: %d\nÉ múltiplo de: %d\n", num1, num2)
	}
	fmt.Println("The last mission is over!")
}
