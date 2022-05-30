//Produtos e-commerce
package main

import "fmt"

type produto struct {
	tipo  string
	nome  string
	preco float64
}

type eletro struct {
	p produto
}

type frios struct {
	p    produto
	temp float64
}

type loja struct {
	produtos []produto
}

type interfProduto interface {
	CalcularCusto() float64
}

func (p produto) CalcularCusto() float64 {
	if p.tipo == "Pequeno" {
		return p.preco
	} else if p.tipo == "Médio" {
		return p.preco + 100
	} else {
		return p.preco + 1000
	}
}

func (e eletro) CalcularCusto() float64 {
	if e.p.tipo == "Brastemp" {
		return e.p.preco
	} else {
		return e.p.preco + 2500
	}
}

type interfEcommerce interface {
	Total() float64
	Adicionar(produto)
}

func (l loja) Total() float64 {
	total := 0.0

	for _, produto := range l.produtos {
		total += produto.preco + produto.CalcularCusto()
	}

	return total
}

func (l *loja) Adicionar(p produto) {
	l.produtos = append(l.produtos, p)
}

func novoProduto(tipo string, nome string, preco float64) produto {
	return produto{tipo, nome, preco}
}

func novaLoja() interfEcommerce {
	return &loja{}
}

func main() {
	// casasBahia := novaLoja()

	// geladeiraBrastemp := eletro{
	// 	produto{
	// 		"Grande",
	// 		"Geladeira Brastemp",
	// 		5500.00,
	// 	},
	// }

	// casasBahia.Adicionar(geladeiraEletrolux)
	// casasBahia.Adicionar(geladeiraBrastemp)

	// fmt.Println(casasBahia.Total())

	//Teste SLICE
	produtos := []string{"TV", "Cel", "Xpto"}

	for _, produto := range produtos {

		fmt.Println("\nProduto ", produto)
	}
	//If Declaracao Curta
	if teste := 5; teste == 5 {
		fmt.Println("Declaracao Curtaß")
	}
}
