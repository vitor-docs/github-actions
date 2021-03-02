package main

import "testing"

func TestSoma(t *testing.T) {
	result := 20
	total := Soma(10, 10)

	if total != result {
		t.Errorf("Resultado da soma é invalido: Resultado %d. Esperado: %d", total, result)
	}
}
