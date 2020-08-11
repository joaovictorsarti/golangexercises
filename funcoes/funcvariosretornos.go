package main

import (
	"fmt"
)

func soma(x ...int) (int, int) {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x)

}

func main() {
	total, quantos := soma(10, 20, 30, 40, 50, 60, 70, 80)
	fmt.Println(total, quantos)
}
