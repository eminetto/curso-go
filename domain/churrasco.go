package domain

import "fmt"

//CalculaChurrasco faz o cálculo do churrasco
func CalculaChurrasco(homens int, mulheres int, criancas int, acompanhamentos bool) {
	tc := mulheres*150 + homens*200 + criancas*100
	tp := mulheres + homens + criancas
	ta := tp * 50
	naoAlcoolicas := 400 * tp
	alcoolicas := 500 * (mulheres + homens)

	fmt.Printf("Total de pessoas: %d\n", tp)
	fmt.Printf("Total de carne (g): %d\n", tc)
	fmt.Printf("Total de acompanhamentos (g): %d\n", ta)
	fmt.Printf("Total de bebidas não alcoolicas (ml) %d\n", naoAlcoolicas)
	fmt.Printf("Total de bebidas alcóolicas (ml) %d\n", alcoolicas)
}
