package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) demonstrar() {
	fmt.Println("o nome completo é:", p.nome, p.sobrenome, "essa pessoa tem", p.idade, "anos")
}

func main() {
	p1 := pessoa{
		nome:      "João",
		sobrenome: "Sarti",
		idade:     21,
	}

	p1.demonstrar()

}
