package infraestructura

import (
	"fmt"
	"log"
	"microConsultaPuntos/adaptador/configuracion"
	"microConsultaPuntos/adaptador/repositorio"
	"microConsultaPuntos/dominio/eventos"
	repoDominio "microConsultaPuntos/dominio/repositorio"
	"microConsultaPuntos/dominio/servicios"
	events "microConsultaPuntos/infraestructura/brokers"
	"microConsultaPuntos/infraestructura/migraciones"
	"net/http"
)

func Run() (err error) {

	configuracion.InitLogs()

	db := repositorio.StartDynamo()
	err = migraciones.CrearTablaUsuario(db)
	if err != nil {
		fmt.Println("Error al intentar crear o actualizar tablas", err)
	}
	usuarioRepositorio := repoDominio.NewUsuarioRepositorio(db)
	repoDominio.SetSearchRepository(usuarioRepositorio)
	defer repoDominio.Close()

	n, err := events.NewNats(fmt.Sprintf("nats://%s", "localhost:4222"))
	if err != nil {
		return err
	}

	err = n.OnCreatedFeed(servicios.OnCreatedFeed)
	if err != nil {
		log.Fatal(err)
	}

	eventos.SetEventStore(n)
	defer eventos.Close()

	router := newRouter()
	if err := http.ListenAndServe(":8874", router); err != nil {
		log.Fatal(err)
		return err
	}
	return
}
