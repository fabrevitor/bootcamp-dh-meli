package main

import "fmt"

func main() {
	calcularMedia(7, 8, 7, 6, 5, 4, 3, 7, 10)
}

func calcularMedia(notas ...int) {
	soma := 0

	for _, n := range notas {
		soma += n
	}

	fmt.Println("Média: ", soma/len(notas))
}
