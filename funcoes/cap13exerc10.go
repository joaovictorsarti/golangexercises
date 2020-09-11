package main

import (
	"fmt"
)

func main() {
	recebefuncao(argumento)

}

func argumento() {
	fmt.Println("Serei argumento")
}

func recebefuncao(x func()) {
	fmt.Println("Atenção")
	x()
}
