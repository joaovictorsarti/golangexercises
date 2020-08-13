package main

import (
	"fmt"
)

func main() {
	x := 10

	func(x int) {
		fmt.Println(x, "vezes 15 é:")
		fmt.Println(x * 15)
	}(x)
}
