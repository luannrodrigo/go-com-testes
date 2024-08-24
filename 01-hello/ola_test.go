package main

import "testing"

func TestOla(t *testing.T) {
	verificaMensagem := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("deve retornar uma saudação com o nome da pessoa", func(t *testing.T) {
		resultado := Ola("Luann", "")
		esperado := "Olá, Luann"

		verificaMensagem(t, resultado, esperado)
	})

	t.Run("deve retornar uma saudação genérica caso não tenha passado uma pessoa", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, Mundo"

		verificaMensagem(t, resultado, esperado)
	})

	t.Run("deve retornar uma saudação em espanhol", func(t *testing.T) {
		resultado := Ola("Luann", "es")
		esperado := "Holla, Luann"

		verificaMensagem(t, resultado, esperado)
	})
}
