package model

import (
	"Go-Basico/loja/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectarComBancoDeDados()
	selectTodosOsProdutos, err := db.Query("select * from produto order by id asc")

	p := Produto{}
	produtos := []Produto{}

	for selectTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
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

func CriarNovoProduto(nome string, descricao string, preco float64, quantidade int) {

	db := db.ConectarComBancoDeDados()
	insereDadosNoBanco, err := db.Prepare("insert into produto (nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func DeletarProduto(id string) {
	db := db.ConectarComBancoDeDados()
	deletaProduto, err := db.Prepare("delete from produto where id = $1")

	if err != nil {
		panic(err.Error())
	}

	deletaProduto.Exec(id)

	defer db.Close()
}

func BuscarProduto(id string) Produto {
	db := db.ConectarComBancoDeDados()

	produtoDoBanco, err := db.Query("select * from produto where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	produtoAtualizado := Produto{}

	for produtoDoBanco.Next() {
		var nome, descricao string
		var id, quantidade int
		var preco float64

		err := produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		produtoAtualizado.Id = id
		produtoAtualizado.Nome = nome
		produtoAtualizado.Descricao = descricao
		produtoAtualizado.Preco = preco
		produtoAtualizado.Quantidade = quantidade
	}

	defer db.Close()
	return produtoAtualizado
}

func AtualizarProduto(id int, nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectarComBancoDeDados()

	atualizaProduto, err := db.Prepare("update produto set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}

	atualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}

func conectarComBancoDeDados() {
	panic("unimplemented")
}
