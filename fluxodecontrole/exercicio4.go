package main

import "fmt"

//Loop for condition

func main() {

	anonascimento := 1998
	anoquequerocontar := 2020

	for {
		if anonascimento > anoquequerocontar {
			break
		}
		fmt.Println(anonascimento)
		anonascimento++
	}

}
