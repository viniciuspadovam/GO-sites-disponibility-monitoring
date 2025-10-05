package main

import "fmt"

func main() {
	fmt.Println("Olá, o que deseja fazer?")

	fmt.Println("1 - Inicial Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var comando int
	fmt.Scan(&comando)

	fmt.Println("O comando selecionado foi", comando)
}
