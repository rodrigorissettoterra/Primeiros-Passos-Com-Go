package main

//Importando pacotes / biliotecas
import (
	"fmt"
	"time"
)

// Definindo valores de constantes
const alongamento = 1
const aquecimento = 5
const corridaLeve = 10
const corridaForte = 10

func main() {
	//Exibindo mensagens com delay dependendo de constantes
	fmt.Println("Período de alongamento...")
	time.Sleep(alongamento * time.Minute)

	fmt.Println("Período de aquecimento...")
	time.Sleep(aquecimento * time.Minute)

	fmt.Println("Período de corrida leve...")
	time.Sleep(corridaLeve * time.Minute)

	fmt.Println("Período de corrida forte...")
	time.Sleep(corridaForte * time.Minute)

	fmt.Println("Período de alongamento...")
	time.Sleep(alongamento * time.Minute)
}
