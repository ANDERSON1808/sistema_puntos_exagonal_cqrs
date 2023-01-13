package infraestructura

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"microMutationPuntos/adaptador"
	"microMutationPuntos/adaptador/configuracion"
	"microMutationPuntos/dominio/servicios"
)

type Server struct {
	configApp        configuracion.Config
	engine           *fiber.App
	servicioPuntos   *servicios.ServicioPuntos
	ServicioUsuarios *servicios.ServicioUsuarios
}

func New(cfg configuracion.Config, servicioPuntos *servicios.ServicioPuntos, servicioUsuarios *servicios.ServicioUsuarios) Server {
	svr := Server{
		configApp:        cfg,
		engine:           fiber.New(),
		servicioPuntos:   servicioPuntos,
		ServicioUsuarios: servicioUsuarios,
	}
	svr.engine = fiber.New(fiber.Config{
		ErrorHandler: adaptador.ErrorHandler,
	})

	svr.registerMiddlewares()
	svr.registerRoutes()
	return svr
}

func (s *Server) Run() error {
	log.Info("Starting the application...")
	return s.engine.Listen(":34963")
}
