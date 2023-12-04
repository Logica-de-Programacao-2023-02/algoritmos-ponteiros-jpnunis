package main

import "fmt"

func verificarParOuImpar(ptr *int) {
	if *ptr%2 == 0 {
		*ptr = 0
	} else {
		*ptr = 1
	}
}

func main() {
	numero := 7

	verificarParOuImpar(&numero)

	if numero == 0 {
		fmt.Println("O número é par.")
	} else {
		fmt.Println("O número é ímpar.")
	}
}
