package aniversario

import (
	"testing"

	"github.com/eminetto/curso-go/domain"
	"github.com/stretchr/testify/assert"
)

func TestCalculaAniversario(t *testing.T) {
	aniversario := NewAniversario()

	resultadoCalculado, _ := aniversario.Calcula(domain.Parametros{
		Homens:          1,
		Mulheres:        1,
		Criancas:        1,
		Acompanhamentos: false,
	})

	resultadoEsperado := Aniversario{
		TotalCarne:           450,
		TotalPessoas:         3,
		TotalAcompanhamentos: 150,
		NaoAlcoolicas:        1200,
		Alcoolicas:           1000,
		TotalPresente:		  300,
	}

	assert.Equal(t, resultadoEsperado, resultadoCalculado)
}

func TestFestaSemPessoas(t *testing.T) {
	aniversario := NewAniversario()

	_, err := aniversario.Calcula(domain.Parametros{
		Homens:          0,
		Mulheres:        0,
		Criancas:        0,
		Acompanhamentos: true,
	})

	assert.Equal(t, err.Error(), "Festa de aniversario só com aniversariante?")
}