package main

import (
	"fmt"
)

func main() {
	argumento("noite")
}

func argumento(s string) {
	if s == "manhã" {
		fmt.Println("Oie bom dia")
	} else if s == "tarde" {
		fmt.Println("Oi boa tarde")
	} else {
		fmt.Println("boa noite")
	}
}
