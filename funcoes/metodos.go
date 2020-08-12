package main

import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
}

func (p pessoa) oibomdia() {
	fmt.Println(p.nome, "diz bom dia")
}

func main() {
	mauricio := pessoa{"mauricio", 30}
	mauricio.oibomdia()
}

//func (receiver) identifier(parameters (returns {
//	code
//}

//o metodo oibomdia só vai ser chamado pelo tipo p pessoa

//definindo algo especifico para o tipo
