package main

import (
	"fmt"
)

func main() {
	x := 10

	y := func(x int) int {
		return x * 15
	}
	fmt.Println(x, "vezes 15 e:", y(x))
}
