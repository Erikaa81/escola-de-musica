package entities

import (
	"errors"
	"strings"
)

type Violao struct {
	Marca     string
	Modelo    string
	TipoCorda string
	Cor       string
	Eletrico  bool
}

var tiposCorda = []string{"AÇO", "NYLON"}

func TipoCordaValidado(tipoCorda string) bool {
	tipoCorda = strings.ToUpper(tipoCorda)
	for _, tipoCordalValidado := range tiposCorda {
		if tipoCordalValidado == tipoCorda {
			return true
		}
	}
	return false
}
func CadastrarViolao(marca, modelo, tipoCorda, cor string, elétrico bool) (Violao, error) {
	if !TipoCordaValidado(tipoCorda) {

		return Violao{}, errors.New("a Corda do violao precisa ser de nylon ou aço")
	}

	return Violao{
		Marca:     marca,
		Modelo:    modelo,
		TipoCorda: tipoCorda,
		Cor:       cor,
		Eletrico:  elétrico,
	}, nil
}
