package main

import "fmt"

func main() {

	var entrada, centena, dezena, unidade, soma int

	fmt.Print("Insira um valor de 3 digitos: ")
	fmt.Scanln(&entrada)

	centena = (entrada / 100)
	dezena = ((entrada / 10) % 10)
	unidade = (entrada % 10)

	soma = ((centena + dezena) + unidade)

	if (unidade % 2) == 0 {
		fmt.Printf("Numero %d eh um numero par\nA soma de seus algarismos resulta em: %d\n", entrada, soma)
	} else {
		fmt.Printf("Numero %d eh um numero impar\nA soma de seus algarismos resulta em: %d\n", entrada, soma)
	}
	fmt.Println("Bye bye")
}
