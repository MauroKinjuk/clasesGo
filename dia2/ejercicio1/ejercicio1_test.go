package ejercicio1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImpuestos(t *testing.T) {

	t.Run("Impuesto salario > 150k", func(t *testing.T) {
		var salario float64 = 160000
		impEsperado := 27

		resultado := CalculoImpuesto(salario)

		assert.Equal(t, impEsperado, resultado, "Impuesto incorrecto para salario de %.2f", salario)
	})

	t.Run("Impuesto salario > 50k", func(t *testing.T) {
		var salario float64 = 60000
		impEsperado := 17

		resultado := CalculoImpuesto(salario)

		assert.Equal(t, impEsperado, resultado, "Impuesto incorrecto para salario de %.2f", salario)
	})

	t.Run("Impuesto salario < 50k", func(t *testing.T) {
		var salario float64 = 49000
		impEsperado := 0

		resultado := CalculoImpuesto(salario)

		assert.Equal(t, impEsperado, resultado, "Impuesto incorrecto para salario de %.2f", salario)
	})
}
