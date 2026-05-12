package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tmpl := template.Must(template.ParseFiles("layout.tmpl"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "current",
			Todos: []Todo{
				{Title: "https://go.dev/ref/spec", Done: true},
				{Title: "https://go.dev/doc/code", Done: true},
				{Title: "https://go.dev/doc/effective-go", Done: true},
				{Title: "https://gobyexample.com", Done: true},
				{Title: "https://gowebexamples.com", Done: false},
			},
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":8080", nil)
}
