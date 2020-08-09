package main

import "fmt"

//exercicio 8,9e 10 criando um map, adicionando um map adicional, fazendo um range pra pegar todos
// e deletando um map
func main() {

	mepe := map[string][]string{
		"eleven": []string{
			"desaparecer", "ser assustadora", "raspar o cabelo",
		},
		"senna": []string{
			"correr", "ser rico", "paulista",
		},
		"robepyke": []string{
			"criar linguagens", "usar roupas diferentes",
		},
		"joao": []string{
			"viajar", "estudar TI", "fazer exercicio", "basquete",
		},
	}

	mepe["kiko"] = []string{"lavar a guitarra", "fogo na guitarra", "viajar"}

	delete(mepe, "kiko")

	for k, v := range mepe {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, " - ", h)

		}
	}

}
