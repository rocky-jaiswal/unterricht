package controllers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to the home page!</h1>")
}
