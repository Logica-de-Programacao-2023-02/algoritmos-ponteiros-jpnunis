package main

import "fmt"

func somaPrimeirosNaturais(ptr *int, n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	*ptr = sum
}

func main() {
	var resultado int
	n := 5

	somaPrimeirosNaturais(&resultado, n)

	fmt.Printf("A soma dos primeiros %d números naturais é: %d\n", n, resultado)
}
