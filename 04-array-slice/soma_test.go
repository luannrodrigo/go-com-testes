package main

import (
	"reflect"
	"testing"
)

func TestSoma(t *testing.T) {
	verificaSomas := func(t *testing.T, resultado, esperado []int) {
		t.Helper()
		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("resultado %v, esperado %v", resultado, esperado)
		}
	}
	t.Run("Deve fazer uma soma total de itens usando array", func(t *testing.T) {
		numeros := [5]int{1, 2, 3, 4, 5}
		resultado := Soma(numeros)
		esperado := 15

		if resultado != esperado {
			t.Errorf("resultado %d, esperado %d, dado %v", resultado, esperado, numeros)
		}
	})

	t.Run("Deve fazer uma soma total de itens usando slice", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4, 5}
		resultado := SomaSlice(numeros)
		esperado := 15

		if resultado != esperado {
			t.Errorf("resultado %d, esperado %d, dado %v", resultado, esperado, numeros)
		}
	})

	t.Run("Deve somar tudo", func(t *testing.T) {
		resultado := SomaTudo([]int{3, 4}, []int{3, 2})
		esperado := []int{7, 5}

		verificaSomas(t, resultado, esperado)
	})

	t.Run("Deve somar todo o resto e não sei mais o porque", func(t *testing.T) {
		resultado := SomaTodoResto([]int{3, 4}, []int{3, 2})
		esperado := []int{4, 2}

		verificaSomas(t, resultado, esperado)
	})

	t.Run("Deve somar os slices de forma segura", func(t *testing.T) {
		resultado := SomaTodoResto([]int{}, []int{3, 2, 4, 4})
		esperado := []int{0, 10}

		verificaSomas(t, resultado, esperado)
	})

}
