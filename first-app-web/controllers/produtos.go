package controllers

import (
	"first-app-web/models"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var templates = template.Must(template.ParseGlob("./templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaTodosOsProdutos()

	templates.ExecuteTemplate(w, "Index", produtos)
	fmt.Println("Fechando conexao com o db...")
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("erro na converção do preço")
		}
		quantidadeConvertido, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("erro na converção da quantidade")
		}

		models.CriarNovoProduto(nome, descricao, precoConvertido, quantidadeConvertido)
	}

	http.Redirect(w, r, "/", 301)
}