package main

import "fmt"

func main() {

	var velocidade_max_permitida, velocidade_do_motorista int
	var nome_do_motorista string

	fmt.Printf("Insira o nome do(a) motorista: ")
	fmt.Scanln(&nome_do_motorista)
	fmt.Printf("Insira a velocidade máxima permitida: ")
	fmt.Scanln(&velocidade_max_permitida)
	fmt.Printf("Insira a velocidade do(a) motorista: ")
	fmt.Scanln(&velocidade_do_motorista)

	if (velocidade_do_motorista >= (velocidade_max_permitida + 1)) && (velocidade_do_motorista <= (velocidade_max_permitida + 10)) {
		fmt.Printf("Motorista: %s\n", nome_do_motorista)
		fmt.Println("Multa: R$50,00")
	} else if (velocidade_do_motorista >= (velocidade_max_permitida + 11)) && (velocidade_do_motorista <= velocidade_max_permitida+30) {
		fmt.Printf("Motorista: %s\n", nome_do_motorista)
		fmt.Println("Multa: R$100,00")
	} else if velocidade_do_motorista >= (velocidade_max_permitida + 31) {
		fmt.Printf("Motista: %s\n", nome_do_motorista)
		fmt.Println("Multa de: R$200,00")
	} else {
		fmt.Printf("Motorista: %s\n", nome_do_motorista)
		fmt.Println("Sem Multa")
	}
}
