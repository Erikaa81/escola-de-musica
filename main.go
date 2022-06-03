package main

import (
	"fmt"

	"github.com/Erikaa81/escola-de-musica/entities"
)

func main() {

	violao, err := entities.CadastrarViolao("", "Classico", "nylon", "preto", true)
	if err != nil {
		fmt.Printf("Erro ao cadastrar violão: %v", err)
	} else {
		fmt.Println("Marca: ", violao.Marca)
		fmt.Println("Tipo de corda do violão: ", violao.TipoCorda)

	}
}


