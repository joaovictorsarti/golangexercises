package main

import "fmt"

//acione else e else

func main() {
	x := 50

	if x > 10 {
		fmt.Println("é maior que x")
	} else if x < 10 {
		fmt.Println("Esse numero é menor que x")
	} else {
		fmt.Println("é igual a x")
	}

}
