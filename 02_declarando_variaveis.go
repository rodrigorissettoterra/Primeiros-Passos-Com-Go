package main

//Importando pacotes / biliotecas
import "fmt"

func main() {
	//Declarando variáveis
	var nome string = "Rodrigo"
	var idade int = 39
	var versao float32 = 1.1

	//Exibindo mensagens concatenadas com variáveis
	fmt.Println("Olá,", nome, ".")
	fmt.Println("Até a data de 15/04/2023, sua idade é de:", idade, "anos.")
	fmt.Println("Esta é a versão:", versao, "do meu primeiro programa feito com a linguagem Go!")
}
