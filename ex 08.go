package main

import "fmt"

func modificarValorComLiberacaoDeMemoria(ptr *int) {
	if ptr == nil {
		return
	}

	*ptr = 42
	ptr = nil
}

func main() {
	numero := new(int)

	modificarValorComLiberacaoDeMemoria(numero)

	if numero == nil {
		fmt.Println("A memória foi liberada corretamente.")
	} else {
		fmt.Println("Erro: a memória não foi liberada.")
	}
}
