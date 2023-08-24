package rota

import "github.com/gorilla/mux"

func Gerar() *mux.Router{
	r := mux.NewRouter()
	
	return Configurar(r)
}