package main

import "fmt"

func main() {
	fmt.Println("Ola!")
	exercicio1905Tarde1()
}

func exercicio1905Tarde1() {
	palavra := "Bootcamp"
	fmt.Println("Tamanho: ", len(palavra))

	for _, r := range palavra {
		fmt.Println(string(r))
	}

}
