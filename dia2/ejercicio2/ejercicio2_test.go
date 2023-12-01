package ejercicio2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromedio(t *testing.T) {
	t.Run("Testea que el promedio", func(t *testing.T) {
		promEsperado := 3.5
		notas := []int{1, 2, 3, 4, 5, 6}
		resultado := PromedioNotas(notas...)

		assert.Equal(t, promEsperado, resultado, "error: el prom esperado es: %2.f", promEsperado)
	})

	t.Run("Test numeros negativos", func(t *testing.T) {
		promEsperado := 0.0
		notas := []int{1, 2, 3, -4, 5, 6}
		resultado := PromedioNotas(notas...)

		assert.Equal(t, promEsperado, resultado, "error: el prom esperado es: %2.f", promEsperado)
	})
}
