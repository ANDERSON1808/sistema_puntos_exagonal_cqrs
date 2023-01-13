package servicios

import (
	"errors"
	"fmt"
	"microMutationPuntos/dominio/entidades"
	"microMutationPuntos/dominio/eventos"
	"microMutationPuntos/dominio/repositorio"
	"microMutationPuntos/infraestructura/Modelos"
	"strconv"
)

type ServicioPuntos struct {
	repositorio *repositorio.PuntosRepositorio
}

func NewServicioPuntos(repositorio *repositorio.PuntosRepositorio) *ServicioPuntos {
	return &ServicioPuntos{repositorio: repositorio}
}
func (p ServicioPuntos) ServicioAcumularPuntos(usuarios *ServicioUsuarios, modelo *Modelos.RequestAcumularPuntos) (err error) {
	puntoExiste, err := p.repositorio.BuscarPuntoId(modelo.PuntoId)
	if err != nil {
		fmt.Println("Error buscar punto id", err)
		return err
	}
	if puntoExiste.DetalleMovimiento != "" {
		return errors.New("el idPunto ya existe")
	}
	puntos := entidades.NewPuntos(modelo.PuntoId, strconv.Itoa(modelo.UsuarioId), strconv.Itoa(modelo.Punto), modelo.DetalleMovimiento)
	usuario := entidades.NewUsuario(modelo.UsuarioId, modelo.NombreUsuario, strconv.Itoa(modelo.Punto))
	err = p.repositorio.AcumularPuntos(puntos)
	if err != nil {
		fmt.Println("Error en AcumularPuntos", err)
		return
	}
	eventos.EmitirActualizationPuntosPorUsuario(usuario)

	return
}
