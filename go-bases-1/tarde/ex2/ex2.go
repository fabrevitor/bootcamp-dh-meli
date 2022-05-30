package main

import "fmt"

func main() {
	clientes := []map[string]interface{}{
		{
			"Nome":          "Vitor",
			"Idade":         24,
			"AnosAtividade": 3,
			"Salario":       500000,
		},
		{
			"Nome":          "João",
			"Idade":         20,
			"AnosAtividade": 1,
			"Salario":       25000,
		},
		{
			"Nome":          "Neymar",
			"Idade":         30,
			"AnosAtividade": 10,
			"Salario":       3912039,
		},
	}

	for _, c := range clientes {
		nome := c["Nome"]
		idade := c["Idade"].(int)
		anosAtividade := c["AnosAtividade"].(int)
		salario := c["Salario"].(int)
		message := ""

		if idade <= 22 {
			message = "\nIdade insuficiente."
		}

		if anosAtividade <= 1 {
			message += "\nAnos de atividade insuficientes."
		}

		if message == "" {
			message = "\nEmpréstimo disponível "

			if salario > 100000 {
				message += "sem juros."
			} else {
				message += "com juros."
			}
		}

		fmt.Println("Olá ", nome, message)
	}

}
