package controllers

import (
	"../models"
	"fmt"
	"net/http"
)

func TodoIndex(w http.ResponseWriter, req *http.Request) {
	todos := models.TodosIndex()
	fmt.Println(todos)
	fmt.Fprintf(w, "<h1>Todos</h1>")
}
