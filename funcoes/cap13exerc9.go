package main

import (
	"fmt"
)

func main() {

	x := retornafunc()

	x()

}

func retornafunc() func() {
	return func() {
		fmt.Println("Olha eu sendo chamada aqui")
	}
}
