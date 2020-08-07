package main

import "fmt"

//switch statement string e identificador será favorito

func main() {

	esporteFavorito := "futebol"

	switch esporteFavorito {
	case "futebol":
		fmt.Println("é fut bro, go brazil")
	case "starcraft":
		fmt.Println("é starcraft bro, go corea")
	case "hipismo":
		fmt.Println("upi upi pocotó")
	}

}
