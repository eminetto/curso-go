package main

import (
	"fmt"
	"log"

	"github.com/eminetto/curso-go/domain"
)

func main() {
	churras, err := domain.CalculaChurrasco(domain.Parametros{
		Homens:          0,
		Mulheres:        2,
		Criancas:        3,
		Acompanhamentos: true,
	})
	if err != nil {
		log.Fatal(err)
		// panic(err)
	}
	imprimeChurrasco(churras)
}

func imprimeChurrasco(c domain.Churrasco) {
	fmt.Printf("Total de pessoas: %d\n", c.TotalPessoas)
	fmt.Printf("Total de carne (g): %d\n", c.TotalCarne)
	fmt.Printf("Total de acompanhamentos (g): %d\n", c.TotalAcompanhamentos)
	fmt.Printf("Total de bebidas não alcoolicas (ml) %d\n", c.NaoAlcoolicas)
	fmt.Printf("Total de bebidas alcóolicas (ml) %d\n", c.Alcoolicas)
}
