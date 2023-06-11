package main

import "testing"

func TestSoma(t *testing.T) {

	soma := soma(15, 15)
	sub := sub(20, 16)

	if soma != 30 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", soma, 30)
	}

	if sub != 5 {
		t.Errorf("Resultado da sub é inválido: Resultado %d. Esperado: %d", sub, 5)
	}
}
