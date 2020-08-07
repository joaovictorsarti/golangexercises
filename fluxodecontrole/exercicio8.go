package main

import "fmt"

//declaração switch, sem switch statement especificado.

func main() {
	x := 0

	switch {
	case x == 0:
		fmt.Println("igual a 0 man")
	case x > 10:
		fmt.Println("Maior que x")
	case x < 10:
		fmt.Println("Menor que x")
	}
}
