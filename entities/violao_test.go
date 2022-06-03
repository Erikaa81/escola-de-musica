package entities

import "testing"

func TestViolao(t *testing.T) {

	violao1 := Violao{Marca: "Giannini", Modelo: "Clássico", TipoCorda: "Nylon", Cor: "preto", Eletrico: true}
	t.Run("Esperado todos os itens de acordo com o cosnstrutor", func(t *testing.T) {
		violaoRecebido, err := CadastrarViolao(violao1.Marca, violao1.Modelo, violao1.TipoCorda, violao1.Cor, violao1.Eletrico)
		violaoEsperado := violao1
		errEsperado := err

		if violaoEsperado.Marca != violaoRecebido.Marca || err != errEsperado {
			t.Errorf(" esperado %v , resultado  %v", violaoEsperado.Marca, violaoRecebido.Marca)
		}

		if violaoEsperado.Modelo != violaoRecebido.Modelo || err != errEsperado {
			t.Errorf(" esperado %v , resultado  %v", violaoEsperado.Modelo, violaoRecebido.Modelo)
		}

		if violaoEsperado.TipoCorda != violaoRecebido.TipoCorda || err != errEsperado {
			t.Errorf(" esperado %v , resultado  %v", violaoEsperado.TipoCorda, violaoRecebido.TipoCorda)
		}

		if violaoEsperado.Cor != violaoRecebido.Cor || err != errEsperado {
			t.Errorf(" esperado %v , resultado  %v", violaoEsperado.Cor, violaoRecebido.Cor)
		}

		if violaoEsperado.Eletrico != violaoRecebido.Eletrico || err != errEsperado {
			t.Errorf(" esperado %v , resultado  %v", violaoEsperado.Eletrico, violaoRecebido.Eletrico)
		}

	})

	violao3 := Violao{TipoCorda: "aquela"}
	t.Run("Esperado erro ao cadastrar violao, pois o tipo de corda informado não é valido", func(t *testing.T) {
		tipoDeCordaRecebido := TipoCordaValidado(violao3.TipoCorda)

		tipoDeCordaEsperada := false
		if tipoDeCordaEsperada != tipoDeCordaRecebido {
			t.Errorf(" esperado , %t, resultado  %t", tipoDeCordaEsperada, tipoDeCordaRecebido)
		}

	})
	violao4 := Violao{TipoCorda: "aço"}
	t.Run("Esperado violao cadastrado pois a corda é de aço", func(t *testing.T) {
		tipoDeCordaRecebido := TipoCordaValidado(violao4.TipoCorda)

		tipoDeCordaEsperada := true
		if tipoDeCordaEsperada != tipoDeCordaRecebido {
			t.Errorf(" esperado , %t, resultado  %t", tipoDeCordaEsperada, tipoDeCordaRecebido)
		}

	})
	violao5 := Violao{TipoCorda: "Nylon"}
	t.Run("Esperado violao cadastrado pois a corda é de nylon", func(t *testing.T) {
		tipoDeCordaRecebido := TipoCordaValidado(violao5.TipoCorda)

		tipoDeCordaEsperada := true
		if tipoDeCordaEsperada != tipoDeCordaRecebido {
			t.Errorf(" esperado , %t, resultado  %t", tipoDeCordaEsperada, tipoDeCordaRecebido)
		}

	})

}
