package loops

func Loops(txt string, repeticoes int) string {
	var resultado string

	for i := 0; i < repeticoes; i++ {
		resultado += txt
	}

	return resultado
}
