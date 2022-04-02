package main

import "fmt"

func main() {
	hello("mundo")
	calc(1, 2, 2.5, true)
}

func hello(word string) {
	fmt.Printf("\n\tolá %s\n", word)
}

func calc(a, b int, c float32, d bool) {
	fmt.Println(a + b)
	fmt.Println(c)
	fmt.Println(float64(a) + float64(b) + float64(c))
}
