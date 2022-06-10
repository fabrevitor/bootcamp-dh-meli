package main

import (
	"fmt"
	"time"
)

type aluno struct {
	nome       string
	sobrenome  string
	rg         string
	dtAdmissao time.Time
}

func visualizarAluno(a aluno) {
	fmt.Printf("Aluno: %s %s\nRG: %s\nData de Admissão: %v", a.nome, a.sobrenome, a.rg, a.dtAdmissao)
}

func main() {

	a := aluno{"Vitor", "Fabre", "5512001", time.Now().UTC()}
	visualizarAluno(a)
}
