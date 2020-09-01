package main

import "fmt"

//função que retorna um int, um int e string

func main() {
	fmt.Println(returnint())
	fmt.Println(returnintstring())

}

func returnint() int {
	return 10
}

func returnintstring() (int, string) {
	return 20, "vinte"
}
