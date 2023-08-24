package rota

import (
	"modulo/src/controllers"
	"net/http"
)

var rotaTarefa = []Rota{
	{
		URI:                "/tarefa",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarTarefa,
		RequerAutenticacao: true,
	},
	{
		URI:                "/tarefas",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarTarefas,
		RequerAutenticacao: true,
	},
	{
		URI:                "/tarefa/{tarefaId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarTarefa,
		RequerAutenticacao: true,
	},
	{
		URI:                "/tarefa/{tarefaId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarTarefa,
		RequerAutenticacao: true,
	},
}
