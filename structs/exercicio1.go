package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {

	p1 := pessoa{
		nome:      "João",
		sobrenome: "Sarti",
		sabores:   []string{"Coco", "limão"}}

	p2 := pessoa{"Frederico", "Batista", []string{"brigadeiro", "uva"}}

	for _, v := range p1.sabores {
		fmt.Println(v)
	}
	fmt.Println(p1)
	for _, v := range p1.sabores {
		fmt.Println(v)
	}
	fmt.Println(p2)
	for _, v := range p2.sabores {
		fmt.Println("Sabores:", v)
	}
}
