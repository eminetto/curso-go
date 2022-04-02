package churrasco

import (
	"encoding/json"
	"errors"

	"github.com/eminetto/curso-go/domain"
)

//Churrasco churrasco calculado
type Churrasco struct {
	TotalCarne           int `json:"total-carne"`
	TotalPessoas         int `json:"total-pessoas"`
	TotalAcompanhamentos int `json:"total-acompanhamentos"`
	NaoAlcoolicas        int `json:"nao-alcoolicas"`
	Alcoolicas           int `json:"alcoolicas"`
}

func (c Churrasco) ToJSON() ([]byte, error) {
	return json.Marshal(c)
}

//Service define um serviço de churrasco
type Service struct{}

//NewChurrasco cria um novo churrasco
func NewChurrasco() Service {
	return Service{}
}

//Calcula faz o cálculo do churrasco
func (s Service) Calcula(p domain.Parametros) (domain.Resultado, error) {
	churras := Churrasco{}
	if p.Homens == 0 || p.Mulheres == 0 {
		// return churras, fmt.Errorf("Homens e mulheres devem ser maiores que zero")
		return churras, errors.New("Homens e mulheres devem ser maiores que zero")
	}

	churras.TotalCarne = p.Mulheres*150 + p.Homens*200 + p.Criancas*100
	churras.TotalPessoas = p.Mulheres + p.Homens + p.Criancas
	churras.TotalAcompanhamentos = churras.TotalPessoas * 50
	churras.NaoAlcoolicas = 400 * churras.TotalPessoas
	churras.Alcoolicas = 500 * (p.Mulheres + p.Homens)
	return churras, nil
}
