package rota

import (
	"modulo/src/controllers"
	"net/http"
)

var rotaUsuario = []Rota{
	{
		URI:                "/usuario",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
}
