package modelos

type DadosAutenticacao struct {
	IDUsuario string `json:"id_usuario"`
	Token     string `json:"token"`
}
