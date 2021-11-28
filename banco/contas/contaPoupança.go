package contas

import "banco/clientes"

type ContaPoupança struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroBanco   int
	Operação      int
	saldo         float64
}

func (c *ContaPoupança) Saque(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaPoupança) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso!", c.saldo
	} else {
		return "Não foi possivel depositar", c.saldo
	}
}

func (c *ContaPoupança) Transferir(valordDaTransf float64, contaParaTransf *ContaPoupança) bool {
	if valordDaTransf < c.saldo && valordDaTransf > 0 {
		c.saldo -= valordDaTransf
		contaParaTransf.Depositar(valordDaTransf)
		return true
	} else {
		return false
	}
}

func (c *ContaPoupança) GetSaldo() float64 {
	return c.saldo
}
