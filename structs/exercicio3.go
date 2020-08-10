package main

import "fmt"

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	tracaoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {
	c1 := sedan{veiculo{4, "vermelho"}, true}
	c2 := caminhonete{
		veiculo: veiculo{
			portas: 4,
			cor:    "roxo",
		},
		tracaoNasQuatro: false,
	}

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c1.modeloLuxo)
	fmt.Println(c2.tracaoNasQuatro)
}
