package rota

import (
	"modulo/src/controllers"
	"net/http"
)

var rotaLogin = []Rota {
	{
		URI: "/login",
		Metodo: http.MethodPost,
		Funcao: controllers.Login,
		RequerAutenticacao: false,
	},
}