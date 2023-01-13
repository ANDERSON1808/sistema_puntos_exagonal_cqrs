package entidades

type Usuario struct {
	UsuarioId     int    `json:"UsuarioId"`
	NombreUsuario string `json:"NombreUsuario"`
	TotalPuntos   string `json:"TotalPuntos"`
}

func NewUsuario(usuarioId int, nombreUsuario string, totalPuntos string) *Usuario {
	return &Usuario{UsuarioId: usuarioId, NombreUsuario: nombreUsuario, TotalPuntos: totalPuntos}
}
