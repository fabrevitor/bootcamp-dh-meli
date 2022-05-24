package main

import "fmt"

func main() {
	palavra := "Teste"

	fmt.Println("Tamanho da palavra: ", len(palavra))

	for _, r := range palavra {
		fmt.Println(string(r))
	}
}
