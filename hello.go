package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Robert"
	idade := 19
	versao := 1.1
	fmt.Println("Olá, sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("--------------------------------")
	fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))
	fmt.Println("O tipo da variável idade é", reflect.TypeOf(idade))
	fmt.Println("O tipo da variável versão é", reflect.TypeOf(versao))
}
