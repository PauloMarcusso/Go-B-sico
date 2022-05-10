package main

import (
	"Go-Basico/loja/route"
	"net/http"
)

func main() {
	route.CarregarRotas()
	http.ListenAndServe(":8000", nil)
}
