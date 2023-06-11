package main

import "testing"

func TestCalculo(t *testing.T) {

	soma := soma(15, 15)
	sub := sub(20, 15)

	if soma != 30 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", soma, 30)
	}

	if sub != 5 {
		t.Errorf("It's a new feature %d. Esperado: %d", sub, 5)
	}
}
