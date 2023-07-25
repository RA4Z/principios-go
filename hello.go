package main

import (
	"fmt"
	//"reflect"
)

func main() {
	nome := "Robert"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)

	if comando == 1 {
		fmt.Println("Monitorando...")

	} else if comando == 2 {
		fmt.Println("Exibindo Logs...")

	} else if comando == 0 {
		fmt.Println("Saindo do Programa...")

	} else {
		fmt.Println("Não conheço este comando")
	}

	// fmt.Println("--------------------------------")
	// fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))
	// fmt.Println("O tipo da variável idade é", reflect.TypeOf(idade))
	// fmt.Println("O tipo da variável versão é", reflect.TypeOf(versao))
}
