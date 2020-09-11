package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado1 float64
}

type circulo struct {
	raio float64
}

func (c circulo) area() {
	result := math.Pi * math.Pow(c.raio, 2)
	fmt.Println("Área do circulo", result)
}

func (q quadrado) area() {
	result := q.lado1 * q.lado1
	fmt.Println("Area do quadrado", result)

}

type info interface {
	area()
}

func medida(i info) {
	i.area()
}

func main() {
	x := quadrado{
		lado1: 15.0,
	}

	y := circulo{
		raio: 5.0,
	}

	//x.area()
	//y.area()

	medida(x)
	medida(y)
}
