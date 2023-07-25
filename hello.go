package main

import (
	"fmt"
	"net/http"
	"os"
	//"reflect"
)

func main() {
	exibeIntroducao()

	for {
		exibeMenu()

		comando := lerComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}

	// if comando == 1 {
	// 	fmt.Println("Monitorando...")

	// } else if comando == 2 {
	// 	fmt.Println("Exibindo Logs...")

	// } else if comando == 0 {
	// 	fmt.Println("Saindo do Programa...")

	// } else {
	// 	fmt.Println("Não conheço este comando")
	// }

	// fmt.Println("--------------------------------")
	// fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))
	// fmt.Println("O tipo da variável idade é", reflect.TypeOf(idade))
	// fmt.Println("O tipo da variável versão é", reflect.TypeOf(versao))
}

func exibeIntroducao() {
	nome := "Robert"
	versao := 1.2
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	var sites [4]string
	sites[0] = "https://www.alura.com.br"
	sites[1] = "https://www.alura.com.br"
	sites[2] = "https://www.caelum.com.br"
	sites[3] = "https://www.alura.com.br"

	fmt.Println(sites)

	site := "https://www.alura.com.br"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}

}

func exibeNomes() {
	nomes := []string{"Douglas", "Robert", "John", "Ethan"}
	nomes = append(nomes, "Joseph")
	fmt.Println(nomes)
}
