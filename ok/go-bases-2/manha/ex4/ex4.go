package main

import (
	"errors"
	"fmt"
)

/*
Exercício 4 - Cálculo de estatísticas

Os professores de uma universidade na Colômbia precisam calcular algumas estatísticas de
notas dos alunos de um curso, sendo necessário calcular os valores mínimo, máximo e médio
de suas notas.
Será necessário criar uma função que indique que tipo de cálculo deve ser realizado (mínimo,
máximo ou média) e que retorna outra função (e uma mensagem caso o cálculo não esteja
definido) que pode ser passado uma quantidade N de inteiros e retorne o cálculo que foi
indicado na função anterior

Ex:

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

minhaFunc, err := operation(minimum)
averageFunc, err := operation(average)
maxFunc, err := operation(maximum)

minValue := minhaFunc(2, 3, 3, 4, 10, 2, 4, 5)
averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
*/

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func calcMin(values ...float64) (float64, error) {
	//TODO -> Implement if len(values) > 0 realizar um for para verificar qual o menor valor do array.
	return 0.0, nil
}

func calcMax(values ...float64) (float64, error) {
	//TODO -> Implement if len(values) > 0 realizar um for para verificar qual o maior valor do array.
	return 0.0, nil
}

func calcAvg(values ...float64) (float64, error) {
	//TODO -> Implement if len(values) > 0 realizar um for para realizar a soma e retornar a divisão pelo len.
	return 0.0, nil
}

func getFunc(typeFunc string) (func(values ...float64) (float64, error), error) {
	switch typeFunc {
	case minimum:
		return calcMin, nil
	case maximum:
		return calcMax, nil
	case average:
		return calcAvg, nil
	default:
		return nil, errors.New("invalid function type")
	}
}

func main() {
	minFunc, _ := getFunc(minimum)
	maxFunc, _ := getFunc(maximum)
	avgFunc, _ := getFunc(average)

	min, _ := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Printf("min: %.2f\n", min)
	avg, _ := avgFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Printf("avg: %.2f\n", avg)
	max, _ := maxFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Printf("max: %.2f\n", max)
}
