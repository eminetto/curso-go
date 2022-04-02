package domain

//Parametros parametros da função CalculaChurrasco
type Parametros struct {
	Homens          int  `json:"homens"`
	Mulheres        int  `json:"mulheres"`
	Criancas        int  `json:"criancas"`
	Acompanhamentos bool `json:"acompanhamentos"`
}

//Festa define o que é uma festa
type Festa interface {
	Calcula(Parametros) (Resultado, error)
}

//Resultado define o que é um resultado
type Resultado interface {
	ToJSON() ([]byte, error)
}
