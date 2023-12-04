package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
	Preco  float64
}

func aplicarDesconto(livroPtr *Livro) {
	if livroPtr != nil {
		livroPtr.Preco *= 0.90
	}
}

func main() {
	livroExemplo := Livro{
		Titulo: "Livro A",
		Autor:  "Autor B",
		Preco:  50.0,
	}

	aplicarDesconto(&livroExemplo)

	fmt.Printf("Novo preço do livro %s: %.2f\n", livroExemplo.Titulo, livroExemplo.Preco)
}
