package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Post struct {
	Id    int
	Title string
	Body  string
}

func main() {
	// Difinir rota
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Iniciando template
		t := template.Must(template.ParseFiles("templates/index.html"))
		// Executar o template
		t.ExecuteTemplate(w, "index.html", nill)
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
