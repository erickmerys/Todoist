package main

import (
	"fmt"
	"log"
	"modulo/src/config"
	rota "modulo/src/router/rotas"
	"net/http"
)

func main() {
	config.Carregar()
	r := rota.Gerar()

	fmt.Printf("Escutando na porta %d", config.Porta)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}