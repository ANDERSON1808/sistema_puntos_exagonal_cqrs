package entidades

type Usuario struct {
	UsuarioId     int    `json:"UsuarioId"`
	NombreUsuario string `json:"NombreUsuario"`
	TotalPuntos   string `json:"TotalPuntos"`
}
