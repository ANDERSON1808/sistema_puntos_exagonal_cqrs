package entidades

type Puntos struct {
	PuntoId           int    `json:"PuntoId"`
	UsuarioId         string `json:"UsuarioId"`
	Punto             string `json:"Punto"`
	DetalleMovimiento string `json:"DetalleMovimiento"`
}

func NewPuntos(puntoId int, usuarioId string, punto string, detalleMovimiento string) *Puntos {
	return &Puntos{PuntoId: puntoId, UsuarioId: usuarioId, Punto: punto, DetalleMovimiento: detalleMovimiento}
}
