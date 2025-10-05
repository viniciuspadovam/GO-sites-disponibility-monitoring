package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Olá, o que deseja fazer?")

	showMenu()

	switch getUserInput() {
	case 1:
		fmt.Println("Monitoramento iniciado.")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa.")
		os.Exit(0)
	default:
		fmt.Println("Comando inválido.")
		os.Exit(-1)
	}
}

func showMenu() {
	fmt.Println("1 - Inicial Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func getUserInput() int {
	var command int
	fmt.Scan(&command)
	return command
}
