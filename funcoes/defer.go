package main

import (
	"fmt"
)

//defer - statement - deixa pra ultima hora

func main() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
}
