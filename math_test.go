package main

import "testing"

func TestSoma(t *testing.T) {
	result := 30
	total := Soma(15, 15)

	if total != result {
		t.Errorf("Resultado da soma é invalido: Resultado %d. Esperado: %d", total, result)
	}
}
