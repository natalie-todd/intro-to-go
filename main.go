package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func main() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/todos", getTodos)
	fmt.Println("Server is running on 8080")
	http.ListenAndServe(":8080", nil)
}

type Todo struct {
	Title string
	Content string
}

type pageVariables struct {
	PageTitle string
	PageTodos []Todo
}

var todos []Todo

func getTodos(w http.ResponseWriter, r *http.Request) {
	pageVariables := pageVariables{
		PageTitle: "Intro To Go",
		PageTodos: todos,
	}
	t, err := template.ParseFiles("todos.html")
	if err != nil {
		fmt.Println(err)
	}
	err = t.Execute(w, pageVariables)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)
}


func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}



