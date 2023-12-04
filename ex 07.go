package main

import "fmt"

type Conta struct {
	Saldo   float64
	Titular string
}

func adicionarValorAoSaldo(contaPtr *Conta, valor float64) {
	contaPtr.Saldo += valor
}

func main() {
	contaExemplo := Conta{
		Saldo:   1000.0,
		Titular: "João",
	}

	valorAAdicionar := 200.0

	adicionarValorAoSaldo(&contaExemplo, valorAAdicionar)

	fmt.Printf("Novo saldo da conta de %s: %.2f\n", contaExemplo.Titular, contaExemplo.Saldo)
}
