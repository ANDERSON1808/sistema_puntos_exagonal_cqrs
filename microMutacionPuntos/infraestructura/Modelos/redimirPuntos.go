package Modelos

type RedimirPuntos struct {
	UsuarioId    int `json:"UsuarioId"`
	PuntoRedimir int `json:"PuntoRedimir"`
}

func (p RedimirPuntos) Validar() (msn []string) {
	if p.PuntoRedimir <= 0 {
		msn = append(msn, "PuntoRedimir no puede ser nulo")
	}
	if p.UsuarioId <= 0 {
		msn = append(msn, "UsuarioId no puede ser nulo")
	}
	return
}
