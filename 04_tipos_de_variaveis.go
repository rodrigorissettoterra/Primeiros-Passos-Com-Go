package main

//Importando pacotes / biliotecas
import (
	"fmt"
	"reflect"
)

func main() {
	//Declarando variáveis
	var nome = "Rodrigo"
	var idade = 39
	var preco = 15.99

	//Exibindo mensagens concatenadas com variáveis e seus tipos
	fmt.Println("--------------------")
	fmt.Println("TIPOS DE VARIÁVEIS")
	fmt.Println("--------------------")
	fmt.Println("O tipo da variável nome é:", reflect.TypeOf(nome))
	fmt.Println("O tipo da variável idade é:", reflect.TypeOf(idade))
	fmt.Println("O tipo da variável preço é:", reflect.TypeOf(preco))
	fmt.Println("--------------------")
}
