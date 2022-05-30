package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

	fmt.Println("Idade do Benjamin: ", employees["Benjamin"])
	fmt.Println("Qtde de funcionários acima de 21 anos: ", verificarIdade(employees, 21))
	fmt.Println("Add Fred and delete Pedro")
	employees["Fred"] = 18
	delete(employees, "Pedro")
	fmt.Println("Qtde de funcionários acima de 21 anos: ", verificarIdade(employees, 21))

}

func verificarIdade(mapa map[string]int, idade int) int {
	qtde := 0
	for _, m := range mapa {
		if m > idade {
			qtde++
		}
	}

	return qtde
}
