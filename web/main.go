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
		if err := t.ExecuteTemplate(w, "index.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
