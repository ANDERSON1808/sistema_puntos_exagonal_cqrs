package servicios

import (
	"fmt"
	"microConsultaPuntos/dominio/entidades"
	"microConsultaPuntos/dominio/repositorio"
	"strconv"
)

func RedimirPuntos(m entidades.RedimirPuntos) {
	fmt.Println("Inicia proceo para redimir puntos usuario")
	usuarioExiste, err := repositorio.BuscarUsuarioId(m.UsuarioId)
	if err != nil {
		fmt.Println("Error al validar si usuario existe", err)
	}
	if usuarioExiste.NombreUsuario != "" {
		actual, err := strconv.Atoi(usuarioExiste.TotalPuntos)
		if err != nil {
			fmt.Println("Error strconv.Atoi", err)
		}
		if actual > 0 {
			if m.PuntoRedimir > actual {
				fmt.Println("los puntos a redimir son superiores a los disponibles")
			} else {
				var total = actual - m.PuntoRedimir
				usuarioExiste.TotalPuntos = strconv.Itoa(total)
				usuarioExiste.UsuarioId = m.UsuarioId
				err = repositorio.ActualizarUsuario(&usuarioExiste)
				if err != nil {
					fmt.Println("Error ActualizarUsuario", err)
				}
				fmt.Println("Finaliza proceso actualizacion usuario")
			}
		} else {
			fmt.Println("El usuario no tiene puntos para redimir")
		}
	}
}
