package main

import (
	"fmt"
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
		fmt.Fprintf(w, "Curso de Golang para Web")
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
