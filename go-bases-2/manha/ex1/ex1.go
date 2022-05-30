package main

import "fmt"

func main() {
	salario1 := 55000.00
	salario2 := 175000.00

	fmt.Println(calcularImposto(salario1))
	fmt.Println(calcularImposto(salario2))

}

func calcularImposto(salario float64) float64 {
	if salario <= 50000.00 {
		return 0.0
	} else if salario <= 150000.00 {
		return salario * 0.17
	}

	return salario * 0.27
}
