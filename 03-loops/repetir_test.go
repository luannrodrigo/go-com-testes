package loops

import (
	"fmt"
	"testing"
)

func TestLoops(t *testing.T) {

	t.Run("Deve repetir a string 2 vezes", func(t *testing.T) {
		resultado := Loops("Repetir", 2)
		esperado := "RepetirRepetir"

		if resultado != esperado {
			t.Errorf("esperado '%s' mas obteve '%s'", esperado, resultado)
		}
	})

	t.Run("Deve repetir permitir que quantidade de repetições seja passada como parametro", func(t *testing.T) {
		resultado := Loops("Repetir", 4)
		esperado := "RepetirRepetirRepetirRepetir"

		if resultado != esperado {
			t.Errorf("esperado '%s' mas obteve '%s'", esperado, resultado)
		}
	})

}

func BenchmarkLoops(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loops("a", 2)
	}
}

func ExampleLoops() {
	fmt.Println(Loops("Exemplo", 3))
}
