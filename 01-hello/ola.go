package main

import "fmt"

const SAUDACAO_PT = "Olá, "
const SAUDACAO_ES = "Holla, "

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	return saudacao(idioma) + nome
}

func saudacao(idioma string) (prefixo string) {
	switch idioma {
	case "es":
		prefixo = SAUDACAO_ES
	default:
		prefixo = SAUDACAO_PT
	}
	return
}

func main() {
	fmt.Println(Ola("Luann", ""))
}
