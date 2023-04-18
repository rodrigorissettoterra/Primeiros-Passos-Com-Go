package main

//Importando pacotes / biliotecas
import "fmt"

func main() {
	//Declarando variáveis
	var precoLeite float32 = 5.99
	var precoOvo float32 = 15.99
	var precoPao float32 = 2.49
	var total float32 = precoLeite + precoOvo + precoPao

	//Exibindo mensagens concatenadas com variáveis e soma
	fmt.Println("--------------------")
	fmt.Println("LISTA DE COMPRAS")
	fmt.Println("--------------------")
	fmt.Println("- Preço do leite: R$", precoLeite, "-> 1 unidade")
	fmt.Println("- Preço do ovo: R$", precoOvo, "-> 1 unidade")
	fmt.Println("- Preço do pão: R$", precoPao, "-> 1 unidade")
	fmt.Println("  TOTAL: R$", total)
	fmt.Println("--------------------")
}
