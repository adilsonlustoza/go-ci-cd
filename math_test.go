package main

import "testing"

func SomaTest(t *testing.T) {

	soma := soma(15, 15)
	if soma != 30 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", soma, 30)
	}

}

func SubTest(t *testing.T) {
	sub := sub(20, 15)
	if sub != 5 {
		t.Errorf("Tá passando, está é uma nova feature %d.  Resultado esperado : %d", sub, 5)
	}

}
