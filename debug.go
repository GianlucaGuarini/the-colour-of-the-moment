package main

import (
	"github.com/the-colour-of-the-moment/api"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Handler)
	http.ListenAndServe(":8080", nil)
}
