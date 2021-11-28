package main

import (
	"first-app-web/routes"
	"fmt"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	fmt.Println("Iniciando servidor na porta :8000")
	http.ListenAndServe(":8000", nil)
}
