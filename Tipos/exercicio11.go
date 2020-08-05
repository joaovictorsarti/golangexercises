package main

import "fmt"

//iota

const (
	_ = 1994 + iota
	b
	c
	d
	e
)

func main() {
	fmt.Println(b, c, d, e)
}
