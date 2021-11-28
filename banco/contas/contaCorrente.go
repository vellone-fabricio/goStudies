package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular    clientes.Titular
	NumAgencia int
	NumConta   int
	saldo      float64
}

func (c *ContaCorrente) Saque(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso!", c.saldo
	} else {
		return "Não foi possivel depositar", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valordDaTransf float64, contaParaTransf *ContaCorrente) bool {
	if valordDaTransf < c.saldo && valordDaTransf > 0 {
		c.saldo -= valordDaTransf
		contaParaTransf.Depositar(valordDaTransf)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) GetSaldo() float64 {
	return c.saldo
}
