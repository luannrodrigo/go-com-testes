package main

func Soma(numeros [5]int) int {
	total := 0

	// for i := 0; i < len(numeros); i++ {
	// 	total += numeros[i]
	// }

	for _, numero := range numeros {
		total += numero
	}
	return total
}

func SomaSlice(numeros []int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}
	return total
}

func SomaTudo(numerosParaSomar ...[]int) (somas []int) {
	quantidadeDeNumeros := len(numerosParaSomar)
	somas = make([]int, quantidadeDeNumeros)

	for i, numeros := range numerosParaSomar {
		somas[i] = SomaSlice(numeros)
	}

	return
}

func SomaTodoResto(numeros ...[]int) []int {
	var somas []int

	for _, numero := range numeros {
		if len(numero) == 0 {
			somas = append(somas, 0)
		} else {
			somas = append(somas, SomaSlice(numero[1:]))
		}
	}
	return somas
}
