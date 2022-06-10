package main

import (
	"errors"
	"fmt"
)

/*
Exercício 5 - Cálculo da quantidade de alimento

Um abrigo de animais precisa descobrir quanta comida comprar para os animais de
estimação. No momento eles só têm tarântulas, hamsters, cachorros e gatos, mas a previsão
é que haja muito mais animais para abrigar.
1. Cães precisam de 10 kg de alimento
2. Gatos 5 kg
3. Hamster 250 gramas.
4. Tarântula 150 gramas.

Precisamos:
1. Implementar uma função Animal que receba como parâmetro um valor do tipo texto
com o animal especificado e que retorne uma função com uma mensagem (caso não
exista o animal)
2. Uma função para cada animal que calcule a quantidade de alimento com base na
quantidade necessária do animal digitado.
Exemplo:

const (
dog = "dog"
cat = "cat"
)

...
animalDog, msg := Animal(dog)
animalCat, msg := Animal(cat)

...
var amount float64
amount+= animaldog(5)
amount+= animalCat(8)
*/

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func dogFunc(quantity int) int {
	return quantity * 10000
}

func catFunc(quantity int) int {
	return quantity * 5000
}

func hamsterFunc(quantity int) int {
	return quantity * 250
}

func tarantulaFunc(quantity int) int {
	return quantity * 150
}

func Animal(typeFunc string) (func(quantity int) int, error) {

	switch typeFunc {
	case dog:
		return dogFunc, nil
	case cat:
		return catFunc, nil
	case hamster:
		return hamsterFunc, nil
	case tarantula:
		return tarantulaFunc, nil
	default:
		return nil, errors.New("invalid func")
	}

}

func main() {
	dfunc, _ := Animal(dog)
	fmt.Printf("quandidade de alimento do(s) cachorro(s) (em gramas): %d gramas", dfunc(5))
	fmt.Println("")

	cfunc, _ := Animal(cat)
	fmt.Printf("quantidade de alimento do(s) gato(s) (em gramas): %d gramas", cfunc(8))
	fmt.Println("")

	_, err := Animal("invalid animal")
	if err != nil {
		fmt.Println("animal inválido")
	}

}
