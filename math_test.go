package main

import "testing"

func TestSoma(t *testing.T) {

	soma := soma(15, 15)

	if soma != 30 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", soma, 30)
	}
}
