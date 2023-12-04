package main

import (
	"fmt"
	"math"
)

func preencherComPrimos(slicePtr *[]int, tamanho int) {
	if slicePtr == nil || tamanho <= 0 {
		return
	}

	numerosPrimos := []int{}
	numeroAtual := 2

	for len(numerosPrimos) < tamanho {
		ehPrimo := true

		for i := 2; i <= int(math.Sqrt(float64(numeroAtual))); i++ {
			if numeroAtual%i == 0 {
				ehPrimo = false
				break
			}
		}

		if ehPrimo {
			numerosPrimos = append(numerosPrimos, numeroAtual)
		}

		numeroAtual++
	}

	*slicePtr = numerosPrimos
}

func main() {
	var primos []int
	tamanhoDesejado := 5

	preencherComPrimos(&primos, tamanhoDesejado)

	fmt.Printf("Os primeiros %d números primos: %v\n", tamanhoDesejado, primos)
}
