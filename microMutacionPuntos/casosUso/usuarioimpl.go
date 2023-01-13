package casosUso

import "microMutationPuntos/dominio/entidades"

type UsuarioImpl interface {
	CrearUsuario(usuario *entidades.Usuario) (err error)
	ActualizarUsuario(usuario *entidades.Usuario) (err error)
	BuscarUsuarioId(usuarioId string) (respuesta entidades.Usuario, err error)
}
