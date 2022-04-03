package aniversario

import (
	"encoding/json"
	"fmt"
	"github.com/eminetto/curso-go/domain"
)

const (
	VALOR_PRESENTE = 100
	QTD_MIN_PESSOAS_NA_FESTA = 2 // Aniversariante e pelo menos mais um né.
)

type Aniversario struct {
	TotalCarne           int `json:"total-carne"`
	TotalPessoas         int `json:"total-pessoas"`
	TotalAcompanhamentos int `json:"total-acompanhamentos"`
	Alcoolicas			 int `json:"alcoolicas"`
	NaoAlcoolicas        int `json:"nao-alcoolicas"`
	TotalPresente		 int `json:"total-presente"`
}

func (a Aniversario) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}

type Service struct{}

func NewAniversario() Service {
	return Service{}
}

/* 
	Calcula calcula os itens para o aniversario
*/
func (s Service) Calcula(p domain.Parametros) (domain.Resultado, error) {
	a := Aniversario{}
	a.TotalPessoas = p.Mulheres + p.Homens + p.Criancas


	if a.TotalPessoas < QTD_MIN_PESSOAS_NA_FESTA {
		return a, fmt.Errorf("Festa de aniversario só com aniversariante?")
	}

	a.TotalCarne = p.Mulheres*150 + p.Homens*200 + p.Criancas*100
	
	a.TotalAcompanhamentos = a.TotalPessoas * 50
	a.NaoAlcoolicas = 400 * a.TotalPessoas
	a.Alcoolicas = 500 * (p.Mulheres + p.Homens)
	a.TotalPresente = VALOR_PRESENTE * a.TotalPessoas

	return a, nil
}
