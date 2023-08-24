package rota

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(w http.ResponseWriter, r *http.Request)
	RequerAutenticacao bool
}


func Configurar(r *mux.Router) *mux.Router{
	rotas := rotaUsuario
	rotas = append(rotas, rotaTarefa...)
	rotas = append(rotas, rotaLogin...)

	for _, rota :=  range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}