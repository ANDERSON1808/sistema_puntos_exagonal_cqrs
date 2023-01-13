package repositorio

import (
	"microConsultaPuntos/dominio/entidades"
)

type UsuariosRepositorio interface {
	Close()
	CrearUsuario(usuario *entidades.Usuario) (err error)
	BuscarUsuarioId(usuarioId int) (respuesta entidades.Usuario, err error)
	ActualizarUsuario(usuario *entidades.Usuario) (err error)
}

var repo UsuariosRepositorio

func SetSearchRepository(r UsuariosRepositorio) {
	repo = r
}

func Close() {
	repo.Close()
}

func CrearUsuario(usuario *entidades.Usuario) (err error) {
	return repo.CrearUsuario(usuario)
}

func BuscarUsuarioId(usuarioId int) (respuesta entidades.Usuario, err error) {
	return repo.BuscarUsuarioId(usuarioId)
}

func ActualizarUsuario(usuario *entidades.Usuario) (err error) {
	return repo.ActualizarUsuario(usuario)
}
