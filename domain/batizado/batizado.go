package batizado

import (
	"encoding/json"
	"fmt"

	"github.com/eminetto/curso-go/domain"
)

type Batizado struct {
	TotalCarne           int `json:"total-carne"`
	TotalPessoas         int `json:"total-pessoas"`
	TotalAcompanhamentos int `json:"total-acompanhamentos"`
	NaoAlcoolicas        int `json:"nao-alcoolicas"`
}

func (b Batizado) ToJSON() ([]byte, error) {
	return json.Marshal(b)
}

type Service struct{}

func NewBatizado() Service {
	return Service{}
}

/* Calcula calcula os itens para o bbatizado
A primeira letra maiuscula no nome da função indica que ela é pública
*/
func (s Service) Calcula(p domain.Parametros) (domain.Resultado, error) {
	b := Batizado{}
	if p.Mulheres <= 0 || p.Homens <= 0 || p.Criancas <= 0 {
		return b, fmt.Errorf("Festa fantasma?")
	}

	b.TotalCarne = p.Mulheres*150 + p.Homens*200 + p.Criancas*100
	b.TotalPessoas = p.Mulheres + p.Homens + p.Criancas
	b.TotalAcompanhamentos = b.TotalPessoas * 50
	b.NaoAlcoolicas = 400 * b.TotalPessoas

	return b, nil
}
