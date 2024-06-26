package main

import (
	"adriano-go/models"
	"html/template"
	"net/http"

	"github.com/alura/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)

}
