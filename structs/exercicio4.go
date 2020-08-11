package main

import (
	"fmt"
)

func main() {
	fmt.Println("")

	x := struct {
		nome      string
		sabor     string
		ondetem   []string
		vaibemcom map[string]string
	}{
		nome:    "Stroopwafel",
		sabor:   "doce",
		ondetem: []string{"na Holanda", "na casa do seu tio rico"},
		vaibemcom: map[string]string{
			"no café":   "chá",
			"no almoço": "cafezinho",
			"na janta":  "de noite não",
		},
	}
	fmt.Println(x)
}
