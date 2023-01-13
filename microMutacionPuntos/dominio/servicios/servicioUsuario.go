package servicios

import (
	"fmt"
	"microMutationPuntos/dominio/entidades"
	"microMutationPuntos/dominio/repositorio"
	"strconv"
)

type ServicioUsuarios struct {
	repositorio *repositorio.UsuarioRepositorio
}

func NewServicioUsuarios(repositorio *repositorio.UsuarioRepositorio) *ServicioUsuarios {
	return &ServicioUsuarios{repositorio: repositorio}
}

func (receiver ServicioUsuarios) ServicioCrearUsuario(usuario *entidades.Usuario) (err error) {
	usuarioExiste, err := receiver.repositorio.BuscarUsuarioId(usuario.UsuarioId)
	if err != nil {
		fmt.Println("Error al validar si usuario existe", err)
		return
	}
	if usuarioExiste.NombreUsuario != "" {
		actual, err := strconv.Atoi(usuarioExiste.TotalPuntos)
		if err != nil {
			return err
		}
		nuevo, err := strconv.Atoi(usuario.TotalPuntos)
		if err != nil {
			return err
		}
		var total = actual + nuevo
		usuario.TotalPuntos = strconv.Itoa(total)
		err = receiver.repositorio.ActualizarUsuario(usuario)
		if err != nil {
			return err
		}
		return err
	}
	err = receiver.repositorio.CrearUsuario(usuario)
	if err != nil {
		fmt.Println("Error repositorio crear usuario", err)
		return
	}
	return
}

func (receiver ServicioUsuarios) ServicioRedimirPuntos(usuario *entidades.Usuario) (err error) {
	err = receiver.repositorio.ActualizarUsuario(usuario)
	if err != nil {
		return err
	}
	return
}
