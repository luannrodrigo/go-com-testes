// type Circulo struct {
// 	raio float64
// }

// func Area(retangulo Retangulo) float64 {
// 	return retangulo.largura * retangulo.altura
// }

package main

import "math"

type Retangulo struct {
	largura float64
	altura  float64
}

func Perimetro(retangulo Retangulo) float64 {
	return 2 * (retangulo.largura + retangulo.altura)
}
func (r Retangulo) Area() float64 {
	return r.largura * r.altura
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

type Triangulo struct {
	Base   float64
	Altura float64
}

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) * 0.5
}
