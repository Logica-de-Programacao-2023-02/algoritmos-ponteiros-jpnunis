package main

import "fmt"

func inverterString(ptr *string) {
	runes := []rune(*ptr)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	*ptr = string(runes)
}

func main() {
	texto := "Hello, World!"

	inverterString(&texto)

	fmt.Println("Texto invertido:", texto)
}
