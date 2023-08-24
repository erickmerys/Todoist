package modelos

import (
	"errors"
	"strings"
	"time"
)

type Tarefa struct {
	ID            uint64    `json:"id,omitempty"`
	Titulo        string    `json:"titulo,omitempty"`
	Descricao     string    `json:"descricao,omitempty"`
	Statu         bool      `json:"statu"`
	TarefaUsuario uint64    `json:"tarefa_usuario,omitempty"`
	CriadaEm      time.Time `json:"CriadaEm,omitempty"`
}

// Preparar vai chamar os metódos para validar e formatar a tarefa recebida
func (tarefa *Tarefa) Preparar() error {
	if erro := tarefa.validar(); erro != nil {
		return erro
	}

	tarefa.formatar()
	return nil
}

func (tarefa *Tarefa) validar() error {
	if tarefa.Titulo == "" {
		errors.New("O título da publicação é obrigatório")
	}
	if tarefa.Descricao == "" {
		errors.New("O conteudo da publicação é obrigatório")
	}

	return nil
}

func (tarefa *Tarefa) formatar() {
	tarefa.Titulo = strings.TrimSpace(tarefa.Titulo)
	tarefa.Descricao = strings.TrimSpace(tarefa.Descricao)
}