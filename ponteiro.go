package main

import "fmt"

func main() {

	x := 11
	//recebevalor(x)
	recebeponteiro(&x)

	fmt.Println(x)

}

func recebevalor(x int) {
	x++
	fmt.Println("Na função: ", x)
}

func recebeponteiro(x *int) {
	*x++
	fmt.Println("Na função: ", *x)

}
