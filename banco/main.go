package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Saque(valorDoBoleto)
}

type verificarConta interface {
	Saque(valor float64) string
}

func main() {
	myAccount := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "sid",
			CPF:       "123.123.42.124-12",
			Profissao: "Bistequeiro",
		},
	}
	otherAccount := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Jacare"}}
	myAccount.Depositar(500)
	fmt.Println("Minha conta", myAccount.GetSaldo(), "Outra conta", otherAccount.GetSaldo())
	myAccount.Transferir(300, &otherAccount)
	fmt.Println("Minha conta", myAccount.GetSaldo(), "Outra conta", otherAccount.GetSaldo())
	PagarBoleto(&myAccount, 149.2)
	fmt.Println("Minha conta", myAccount.GetSaldo(), "Outra conta", otherAccount.GetSaldo())
}
