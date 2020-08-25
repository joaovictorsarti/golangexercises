package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type dentista struct {
	pessoa
	dentestirados int
	salario       float64
}

type arquiteto struct {
	pessoa
	typeconstruct string
	nvlloucura    string
}

func (x dentista) oibomdia() {
	fmt.Println("Meu nome é", x.nome, "e eu já arranquei:", x.dentestirados, "e ouve só: BOM DIA!")
}

func (x arquiteto) oibomdia() {
	fmt.Println("Meu nome é", x.nome, "e ouve só: bom dia")
}

type gente interface {
	oibomdia()
}

func serhumano(g gente) {
	g.oibomdia()
	switch g.(type) {
	case dentista:
		fmt.Println("Eu ganho:", g.(dentista).salario)
	case arquiteto:
		fmt.Println("Eu construo:", g.(arquiteto).typeconstruct)
	}
}

func main() {
	mrdente := dentista{
		pessoa: pessoa{
			nome:      "João",
			sobrenome: "Sarti",
			idade:     21,
		},
		dentestirados: 2100,
		salario:       5000,
	}
	mrpredio := arquiteto{
		pessoa: pessoa{
			nome:      "Felipe",
			sobrenome: "Sarti",
			idade:     18,
		},
		typeconstruct: "casas",
		nvlloucura:    "muito louco mesmo",
	}

	mrdente.oibomdia()
	mrpredio.oibomdia()
	fmt.Println("")
	serhumano(mrpredio)
	serhumano(mrdente)
}
