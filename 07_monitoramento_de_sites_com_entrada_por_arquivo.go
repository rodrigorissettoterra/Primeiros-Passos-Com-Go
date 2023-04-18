package main

// Importando pacotes / biliotecas
import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// Definindo valores de constantes
const monitoramentos = 2

// Definindo a função principal e chamando outras funções
func main() {
	exibeIntroducao()
	for {
		exibeMenu()
		comando := leComando()
		switch comando {
		case 1:
			exibeMenuMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

// Definindo as boas-vindas
func exibeIntroducao() {
	var nome string
	fmt.Println("")
	fmt.Print("Digite seu nome: ")
	fmt.Scan(&nome)
	versao := 1.3
	fmt.Println("")
	fmt.Println("Olá,", nome)
	fmt.Println("Este programa está na versão", versao)
}

// Definindo a função para exibir o menu principal
func exibeMenu() {
	fmt.Println("")
	fmt.Println("Selecione a opção desejada")
	fmt.Println("------------------------------")
	fmt.Println("1- Monitoramento de sites")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
	fmt.Println("------------------------------")
	fmt.Print("Sua opção: ")
}

// Definindo a função de exibição do menu secundário ou de monitoramento
func exibeMenuMonitoramento() {
	fmt.Println("")
	fmt.Println("Como deseja enviar o endereço do site a ser verificado?")
	fmt.Println("------------------------------------------------------------")
	fmt.Println("1- Por digitação")
	fmt.Println("2- Por arquivo")
	fmt.Println("0- Retornar para o menu inicial")
	fmt.Println("------------------------------------------------------------")
	fmt.Print("Sua opção: ")

	selecaoMenuMonitoramento := leComando()
	switch selecaoMenuMonitoramento {
	case 1:
		iniciarMonitoramentoDigitado()
	case 2:
		iniciaMonitoramentoArquivo()
	case 0:
		fmt.Println("Retornando ao menu inicial")
		exibeMenu()
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}
}

// Definindo uma função para ler entradas
func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	return comandoLido
}

// Definindo uma função para receber o endereço de sites manualmente
func iniciarMonitoramentoDigitado() {
	fmt.Println("Qual site você gostaria de monitorar?")
	var sites string
	fmt.Scan(&sites)
	testaSite(sites)
}

// Definindo uma função para fazer o monitoramento de endereços de sites, à partir de um arquivo .txt
func iniciaMonitoramentoArquivo() {
	sites := lerSitesArquivo()
	for i, site := range sites {
		fmt.Println("Testando site", i, ":", site)
		testaSite(site)
	}
}

// Definindo uma função para monitorar sites
func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O site informado:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("O site informado:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
	fmt.Println("")
}

// Definindo uma função para fazer a leitura do arquivo .txt
func lerSitesArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}

// Definindo uma função que registra Logs em um arquivo .txt
func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

// Definindo uma função que exibe os Logs já armazenados em um arquivo .txt
func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(arquivo))
}
