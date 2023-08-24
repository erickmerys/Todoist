package repository

import (
	"database/sql"
	"errors"
	"modulo/src/modelos"
)

type Tarefas struct {
	db *sql.DB
}

func NovoRepositorioDeTarefas(db *sql.DB) *Tarefas {
	return &Tarefas{db}
}

func (repositorio Tarefas) Criar(tarefa modelos.Tarefa) (uint64, error) {
	statement, erro := repositorio.db.Prepare("insert into Tarefas(titulo, descricao, statu, tarefa_usuario) values(?,?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(tarefa.Titulo, tarefa.Descricao, tarefa.Statu, tarefa.TarefaUsuario)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (repositorio Tarefas) BuscarTarefa(usuarioID uint64) ([]modelos.Tarefa, error) {
	linhas, erro := repositorio.db.Query(`
		select T.id, T.titulo, T.descricao, T.statu, T.tarefa_usuario, T.CriadaEm from usuarios U
		inner join Tarefas T
		on U.id = T.tarefa_usuario
		where U.id = ?
	`, usuarioID,)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var tarefas []modelos.Tarefa

	for linhas.Next(){
		var tarefa modelos.Tarefa

		if erro = linhas.Scan(
			&tarefa.ID,
			&tarefa.Titulo,
			&tarefa.Descricao,
			&tarefa.Statu,
			&tarefa.TarefaUsuario,
			&tarefa.CriadaEm,
		); erro != nil {
			return nil, erro
		}

		tarefas = append(tarefas, tarefa)
	}

	return tarefas, nil
}

func (repositorio Tarefas) BuscarPorID(tarefaID, usuarioID uint64) (modelos.Tarefa, error) {
	linhas, erro := repositorio.db.Query(`
		select * from Tarefas
		where id = ? and tarefa_usuario = ?
	`, tarefaID, usuarioID)
	if erro != nil {
		return modelos.Tarefa{}, erro
	}
	defer linhas.Close()

	var tarefa modelos.Tarefa

	if linhas.Next() {
		if erro = linhas.Scan(
			&tarefa.ID,
			&tarefa.Titulo,
			&tarefa.Descricao,
			&tarefa.Statu,
			&tarefa.TarefaUsuario,
			&tarefa.CriadaEm,
		); erro != nil {
			return modelos.Tarefa{}, erro
		}
		return tarefa, nil
	}
	return modelos.Tarefa{}, errors.New("Não encontrado ou não autorizado a fazer essa modificação!")

}

func (repositorio Tarefas) AtualizarTarefa(tarefaID uint64, tarefa modelos.Tarefa) error {
	statement, erro := repositorio.db.Prepare("update Tarefas set titulo = ?, descricao = ?, statu = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(tarefa.Titulo, tarefa.Descricao, tarefa.Statu, tarefaID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Tarefas) DeletarTarefa(tarefaID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from Tarefas where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(tarefaID); erro != nil {
		return erro
	}

	return nil
}