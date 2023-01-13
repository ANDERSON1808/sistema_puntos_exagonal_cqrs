package Modelos

type RequestAcumularPuntos struct {
	NombreUsuario     string `json:"NombreUsuario"`
	PuntoId           int    `json:"PuntoId"`
	UsuarioId         int    `json:"UsuarioId"`
	Punto             int    `json:"Punto"`
	DetalleMovimiento string `json:"DetalleMovimiento"`
}

func (p RequestAcumularPuntos) Validar() (errs []string) {
	if p.PuntoId <= 0 {
		errs = append(errs, "PuntoId no puede ser nulo")
	}
	if p.Punto <= 0 {
		errs = append(errs, "punto no puede ser nulo")
	}
	if p.UsuarioId <= 0 {
		errs = append(errs, "UsuarioId no puede ser nulo")
	}
	if p.DetalleMovimiento == "" {
		errs = append(errs, "DetalleMovimiento no puede ser nulo")
	}
	if p.NombreUsuario == "" {
		errs = append(errs, "NombreUsuario no puede ser nulo")
	}
	return errs
}
