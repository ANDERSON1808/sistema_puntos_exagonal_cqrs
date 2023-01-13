package infraestructura

import (
	"fmt"
	"github.com/spf13/viper"
	"microMutationPuntos/adaptador/configuracion"
	"microMutationPuntos/adaptador/repositorio"
	dominioRepositorio "microMutationPuntos/dominio/repositorio"
	"microMutationPuntos/dominio/servicios"
	"microMutationPuntos/infraestructura/migraciones"
)

func Run() (err error) {

	configuracion.InitLogs()

	var cfg configuracion.Config

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return err
	}

	db := repositorio.StartDynamo()
	err = migraciones.CrearTablaUsuario(db)
	err = migraciones.CrearTablaPuntos(db)
	if err != nil {
		fmt.Println("Error al intentar crear o actualizar tablas", err)
	}
	puntosRepositorio := dominioRepositorio.NewPuntosRepositorio(db)
	servicioPuntos := servicios.NewServicioPuntos(puntosRepositorio)

	usuarioRepositorio := dominioRepositorio.NewUsuarioRepositorio(db)
	servicioUsuario := servicios.NewServicioUsuarios(usuarioRepositorio)
	svr := New(cfg, servicioPuntos, servicioUsuario)

	return svr.Run()
}
