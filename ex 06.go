package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
}

func alterarTituloSeAnonimo(livroPtr *Livro) {
	if livroPtr.Autor == "Anônimo" {
		livroPtr.Titulo = "Desconhecido"
	}
}

func main() {
	livroExemplo := Livro{
		Titulo: "Livro A",
		Autor:  "Anônimo",
	}

	alterarTituloSeAnonimo(&livroExemplo)

	fmt.Printf("Novo título do livro: %s\n", livroExemplo.Titulo)
}
