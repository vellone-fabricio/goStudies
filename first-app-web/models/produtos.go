package models

import (
	"first-app-web/db"
	"fmt"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	fmt.Println("Abrindo conexao...")
	db := db.ConnectaComBancoDeDados()
	selectDeTodosOsProdutos, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}
	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()

	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConnectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletaProduto(idDoProduto string) {
	db := db.ConnectaComBancoDeDados()

	deletarOProduto, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarOProduto.Exec(idDoProduto)
	defer db.Close()
}

func EditaProduto(idDoProduto string) Produto {
	db := db.ConnectaComBancoDeDados()

	produtoDoBanco, err := db.Query("select * from produtos where id=$1", idDoProduto)
	if err != nil {
		panic(err.Error())
	}

	produtoParaAtt := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtoParaAtt.Id = id
		produtoParaAtt.Nome = nome
		produtoParaAtt.Descricao = descricao
		produtoParaAtt.Preco = preco
		produtoParaAtt.Quantidade = quantidade

	}
	defer db.Close()

	return produtoParaAtt
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConnectaComBancoDeDados()

	atualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	atualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
