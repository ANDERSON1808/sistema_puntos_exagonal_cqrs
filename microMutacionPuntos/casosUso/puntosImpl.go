package casosUso

import "microMutationPuntos/dominio/entidades"

type PuntosImpl interface {
	AcumularPuntos(puntos *entidades.Puntos) (err error)
	BuscarPuntoId(punto int) (respuesta entidades.Puntos, err error)
}
