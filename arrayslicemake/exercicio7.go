package main

import "fmt"

//array
func main() {
	slice := [][]string{
		[]string{
			"Laura",
			"Rodrigues",
			"Skin care",
		},
		[]string{
			"João",
			"Sarti",
			"TI",
		},
		[]string{
			"fiipe",
			"pontes",
			"dinotrain",
		},
	}
	for _, v := range slice {
		fmt.Println(v[0])
		for _, item := range v {
			fmt.Println("\t", item)
		}
	}
}
