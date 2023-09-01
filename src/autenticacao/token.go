package autenticacao

import (
	"errors"
	"fmt"
	"modulo/src/banco"
	"modulo/src/config"
	"modulo/src/repository"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte(config.SecretKey))
}

// Validar verifica se o token passado na requisiçãoé valido
func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornaChaveDeVerificacao)
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Token inválido")
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func retornaChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}

func ExtrairUsuarioID(r *http.Request) (uint64, error) {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornaChaveDeVerificacao)
	if erro != nil {
		return 0, erro
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioID, erro := strconv.ParseUint(fmt.Sprintf("%0.f", permissoes["usuarioId"]), 10, 64)
		if erro != nil {
			return 0, erro
		}

		db, erro := banco.Conectar()
		if erro != nil {
			errors.New("Erro ao tentar acessar o banco de dados!")
		}
		defer db.Close()

		repositorio := repository.NovoRepositorioAutenticacao(db)
		if erro = repositorio.BuscarToken(fmt.Sprintf("%0.f", permissoes["usuarioId"]), tokenString); erro != nil {
			errors.New("Usuário não autorizado!")
		}
		
		return usuarioID, nil
	}

	return 0, erro
}
