package main

import "fmt"

func main() {
	var ladoX, ladoY, ladoZ int

	fmt.Print("Informe os lados do triângulo: ")
	fmt.Scanln(&ladoX, &ladoY, &ladoZ)

	if ((ladoX + ladoY) > ladoZ) && ((ladoX + ladoZ) > ladoY) && ((ladoY + ladoZ) > ladoX) {
		if (ladoX == ladoY) && (ladoY == ladoZ) {
			fmt.Println("Forma um Triângulo : Equilátero")
		} else if (ladoX != ladoY) && (ladoY != ladoZ) {
			fmt.Println("Forma um Triângulo : Escaleno")
		} else {
			fmt.Println("Forma um Triângulo : Isósceles")
		}
	} else {
		fmt.Println("Não forma um triângulo")
	}
}
