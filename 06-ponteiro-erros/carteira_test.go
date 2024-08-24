package main

import (
	"fmt"
	"testing"
)

func TestCarteira(t *testing.T) {
	confirmaSaldo := func(t *testing.T, carteira Carteira, esperado Bitcoin) {
		t.Helper()
		resultado := carteira.Saldo()

		if resultado != esperado {
			t.Errorf("resultado %s, esperado %s", resultado, esperado)
		}
	}

	t.Run("Deve permitir depositar", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10))

		fmt.Printf("O endereço do saldo no teste é %v \n", &carteira.saldo)

		confirmaSaldo(t, carteira, Bitcoin(10))
	})

	t.Run("Deve permitir retirar", func(t *testing.T) {
		carteira := Carteira{saldo: Bitcoin(200)}
		carteira.Depositar(Bitcoin(10))
		carteira.Depositar(Bitcoin(10))

		carteira.Retirar(Bitcoin(2))

		esperado := Bitcoin(218)

		fmt.Printf("O endereço do saldo no teste é %v \n", &carteira.saldo)

		confirmaSaldo(t, carteira, esperado)
	})

	t.Run("Retirar saldo insuficiente", func(t *testing.T) {
		saldo := Bitcoin(2000)
		carteira := Carteira{saldo: saldo}

		error := carteira.Retirar(Bitcoin(2))

		confirmaSaldo(t, carteira, saldo)

		if error == nil {
			t.Error("Esperava um erro mas nenhum ocorreu")
		}
	})

}
