package main

import "fmt"

func somaDosUltimosDigitos(ptr *int) {
	numero := *ptr
	digito1 := numero % 10
	numero /= 10
	digito2 := numero % 10
	*ptr = digito1 + digito2
}

func main() {
	var valor int = 1234

	somaDosUltimosDigitos(&valor)

	fmt.Println("Novo valor após a soma dos dois últimos dígitos:", valor)
}
