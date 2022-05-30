package main

import "fmt"

const (
	catA = "A"
	catB = "B"
	catC = "C"
)

func main() {
	calcularSalario(catC, 162)
	calcularSalario(catB, 176)
	calcularSalario(catA, 172)
}

func calcularSalario(categoria string, horasTrab float64) {
	salario := 0.0

	switch categoria {
	case catA:
		salario = 1000.00 * horasTrab
	case catB:
		if horasTrab > 160 {
			salario = (1500 * horasTrab) * 1.20
		} else {
			salario = 1500 * horasTrab
		}
	case catC:
		if horasTrab > 160 {
			salario = (3000 * horasTrab) * 1.50
		} else {
			salario = (3000 * horasTrab)
		}
	}
	fmt.Println(salario)
}
