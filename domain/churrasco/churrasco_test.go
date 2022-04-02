package churrasco

import (
	"testing"

	"github.com/eminetto/curso-go/domain"
	"github.com/stretchr/testify/assert"
)

func TestCalculaChurrascoFantasma(t *testing.T) {
	s := NewChurrasco()
	_, err := s.Calcula(domain.Parametros{
		Homens:          0,
		Mulheres:        1,
		Criancas:        2,
		Acompanhamentos: true,
	})
	assert.Equal(t, err.Error(), "Homens e mulheres devem ser maiores que zero")
	// if err.Error() != "Homens e mulheres devem ser maiores que zero" {
	// 	t.Errorf("Recebido %s esperado %s", err.Error(), "Homens e mulheres devem ser maiores que zero")
	// }
}

func TestCalculaChurrascoBombando(t *testing.T) {
	s := NewChurrasco()
	ch, err := s.Calcula(domain.Parametros{
		Homens:          3,
		Mulheres:        3,
		Criancas:        6,
		Acompanhamentos: true,
	})
	assert.Nil(t, err)
	// if err != nil {
	// 	t.Errorf("Recebido %v esperado %v", err, nil)
	// }
	esperado := Churrasco{
		TotalCarne:           1650,
		TotalPessoas:         12,
		TotalAcompanhamentos: 600,
		NaoAlcoolicas:        4800,
		Alcoolicas:           3000,
	}
	assert.Equal(t, ch, esperado)
	// if ch != esperado {
	// 	t.Errorf("Recebido %v esperado %v", ch, esperado)
	// }
}
