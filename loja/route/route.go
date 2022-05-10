package route

import (
	"Go-Basico/loja/controller"
	"net/http"
)

func CarregarRotas() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/new", controller.New)
	http.HandleFunc("/insert", controller.Insert)
}
