package main

//Importando pacotes / biliotecas
import (
	"fmt"
	"reflect"
)

func main() {
	//Declarando variáveis sem 'var'
	nome := "Rodrigo"
	idade := 39
	preco := 15.99

	//Exibindo mensagens concatenadas com variáveis e seus tipos
	fmt.Println("--------------------")
	fmt.Println("TIPOS DE VARIÁVEIS")
	fmt.Println("--------------------")
	fmt.Println("O tipo da variável nome é:", reflect.TypeOf(nome))
	fmt.Println("O tipo da variável idade é:", reflect.TypeOf(idade))
	fmt.Println("O tipo da variável preço é:", reflect.TypeOf(preco))
	fmt.Println("--------------------")
	fmt.Println("Para declarar variáveis sem utilizar 'var', basta incluir ':' antes do sinal '='.")
	fmt.Println("--------------------")
}
