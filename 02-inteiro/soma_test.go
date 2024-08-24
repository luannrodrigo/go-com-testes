package main

import "testing"

func TestSoma(t *testing.T) {
	soma := Soma(2, 2)

	esperado := 4

	if soma != esperado {
		t.Errorf("esperado '%d', resultado '%d'", esperado, soma)
	}

}
