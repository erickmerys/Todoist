package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"modulo/src/autenticacao"
	"modulo/src/banco"
	"modulo/src/modelos"
	"modulo/src/repository"
	"modulo/src/respostas"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CriarTarefa cria uma tarefa e guarda dentro do banco de dados
func CriarTarefa(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var tarefa modelos.Tarefa
	if erro = json.Unmarshal(corpoRequest, &tarefa); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	tarefa.TarefaUsuario = usuarioID

	if erro = tarefa.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDeTarefas(db)
	tarefa.ID, erro = repositorio.Criar(tarefa)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, nil)
}

// BuscarTarefa busca todas as tarefas de um usuário
func BuscarTarefas(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDeTarefas(db)
	tarefas, erro := repositorio.BuscarTarefa(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, tarefas)
}

// AtualizarTarefa atualiza uma tarefa criada pelo usuário
func AtualizarTarefa(w http.ResponseWriter, r *http.Request) {
	UsuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	tarefaID, erro := strconv.ParseUint(parametros["tarefaId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDeTarefas(db)
	_, erro = repositorio.BuscarPorID(tarefaID, UsuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusForbidden, erro)
		return
	}

	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var tarefa modelos.Tarefa

	if erro = json.Unmarshal(corpoRequest, &tarefa); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = tarefa.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repositorio.AtualizarTarefa(tarefaID, tarefa); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// DeletarTarefa deleta uma tarefa criada pelo usuário
func DeletarTarefa(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	tarefaID, erro := strconv.ParseUint(parametros["tarefaId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDeTarefas(db)
	_ , erro = repositorio.BuscarPorID(tarefaID, usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	if erro = repositorio.DeletarTarefa(tarefaID); erro != nil {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível apagar uma tarefa que nã seja a sua!"))
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}
