package main

import "fmt"

func main() {

	var chave string

	fmt.Scanln(&chave)

	if chave == "admin" {
		fmt.Println("$$$ Acesso Liberado, chave {admin} $$$\n Fim do Programa!")
	} else {
		fmt.Println("Fim do Programa!")
	}
}
