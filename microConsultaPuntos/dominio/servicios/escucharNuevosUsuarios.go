package servicios

import (
	"fmt"
	"microConsultaPuntos/dominio/entidades"
	"microConsultaPuntos/dominio/repositorio"
	"strconv"
)

func OnCreatedFeed(m *entidades.Usuario) {
	usuarioExiste, err := repositorio.BuscarUsuarioId(m.UsuarioId)
	if err != nil {
		fmt.Println("Error al validar si usuario existe", err)
	}
	if usuarioExiste.NombreUsuario != "" {
		actual, err := strconv.Atoi(usuarioExiste.TotalPuntos)
		if err != nil {
			fmt.Println("Error strconv.Atoi", err)
		}
		nuevo, err := strconv.Atoi(m.TotalPuntos)
		if err != nil {
			fmt.Println("Error strconv.Atoi", err)
		}
		var total = actual + nuevo
		m.TotalPuntos = strconv.Itoa(total)
		err = repositorio.ActualizarUsuario(m)
		if err != nil {
			fmt.Println("Error ActualizarUsuario", err)
		}
	}
	err = repositorio.CrearUsuario(m)
	if err != nil {
		fmt.Println("Error repositorio crear usuario", err)
	}
}
