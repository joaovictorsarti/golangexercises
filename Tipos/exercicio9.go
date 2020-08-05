package main

import "fmt"

//deslocamento de bits

func main() {

	x := 200

	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
	y := x << 1
	fmt.Printf("%d\t%b\t%#x\n", y, y, y)

}
