package domain

import "errors"

//Parametros parametros da função CalculaChurrasco
type Parametros struct {
	Homens          int
	Mulheres        int
	Criancas        int
	Acompanhamentos bool
}

//Churrasco churrasco calculado
type Churrasco struct {
	TotalCarne           int
	TotalPessoas         int
	TotalAcompanhamentos int
	NaoAlcoolicas        int
	Alcoolicas           int
}

//CalculaChurrasco faz o cálculo do churrasco
func CalculaChurrasco(p Parametros) (Churrasco, error) {
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
