package repository

import (
	"database/sql"
	"modulo/src/modelos"
)

type Token struct {
	db *sql.DB
}

func NovoRepositorioAutenticacao(db *sql.DB) *Token {
	return &Token{db}
}

func (repositorio Token) SalvarToken(Token modelos.DadosAutenticacao) (modelos.DadosAutenticacao, error) {
	statement, erro := repositorio.db.Prepare("insert into Token(id_usuario, token) values(?, ?)")
	if erro != nil {
		return modelos.DadosAutenticacao{}, erro
	}
	defer statement.Close()

	_, erro = statement.Exec(Token.IDUsuario, Token.Token)
	if erro != nil {
		return modelos.DadosAutenticacao{}, erro
	}

	return modelos.DadosAutenticacao{}, erro
}

func (repositorio Token) BuscarToken(usuarioID, token string) error {
	linhas, erro := repositorio.db.Query(`
	select id_usuario, token from Token
	where id_usuario = ? && token = ?
	`, usuarioID, token)
	if erro != nil {
		return erro
	}
	defer linhas.Close()

	if linhas.Next() {
		linhas.Scan(
			&usuarioID,
			&token,
		)
	}

	return erro
}
