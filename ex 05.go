package main

import (
	"fmt"
	"math"
)

func mediaComPi(ptr *float64) {
	const pi = math.Pi
	*ptr = (*ptr + pi) / 2
}

func main() {
	var numero float64 = 3.14

	mediaComPi(&numero)

	fmt.Println("Novo valor após a média com Pi:", numero)
}
