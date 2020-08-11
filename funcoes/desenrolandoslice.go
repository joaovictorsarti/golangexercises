package main

import (
	"fmt"
)

func soma(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma

}

func main() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80}
	total := soma(slice...)
	fmt.Println(total)
}
