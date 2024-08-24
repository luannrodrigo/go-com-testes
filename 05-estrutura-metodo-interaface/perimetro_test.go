package main

import "testing"

func TestPerimetro(t *testing.T) {
	retangulo := Retangulo{10.0, 10.0}
	resultado := Perimetro(retangulo)
	esperado := 40.0

	if resultado != esperado {
		t.Errorf("resultado %2.f esperado %2.f", resultado, esperado)
	}

}

type Forma interface {
	Area() float64
}

func TestArea(t *testing.T) {
	testesArea := []struct {
		nome    string
		forma   Forma
		temArea float64
	}{
		{nome: "Retângulo", forma: Retangulo{largura: 12, altura: 6}, temArea: 72.0},
		{nome: "Círculo", forma: Circulo{Raio: 10}, temArea: 314.1592653589793},
		{nome: "Triângulo", forma: Triangulo{Base: 12, Altura: 6}, temArea: 36.0},
	}

	for _, tt := range testesArea {
		t.Run(tt.nome, func(t *testing.T) {
			resultado := tt.forma.Area()
			if resultado != tt.temArea {
				t.Errorf("%#v resultado %.2f, esperado %.2f", tt.forma, resultado, tt.temArea)
			}
		})
	}
}

// func TestArea(t *testing.T) {
// 	testesArea := []struct {
// 		forma    Forma
// 		esperado float64
// 	}{
// 		{Retangulo{12, 6}, 72.0},
// 		{Circulo{10}, 314.1592653589793},
// 	}

// 	for _, tt := range testesArea {
// 		resultado := tt.forma.Area()
// 		if resultado != tt.esperado {
// 			t.Errorf("resultado %.2f, esperado %.2f", resultado, tt.esperado)
// 		}
// 	}

// 	verificaArea := func(t *testing.T, forma Forma, esperado float64) {
// 		t.Helper()
// 		resultado := forma.Area()

// 		if resultado != esperado {
// 			t.Errorf("resultado %.2f, esperado %.2f", resultado, esperado)
// 		}
// 	}

// 	t.Run("retângulos", func(t *testing.T) {
// 		retangulo := Retangulo{12.0, 6.0}
// 		verificaArea(t, retangulo, 72.0)
// 	})

// 	t.Run("círculos", func(t *testing.T) {
// 		circulo := Circulo{10}
// 		verificaArea(t, circulo, 314.1592653589793)
// 	})
// }
