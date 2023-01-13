package infraestructura

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"microMutationPuntos/infraestructura/Modelos"
)

func (s *Server) RedimirPuntoHandler(ctx *fiber.Ctx) error {
	return ctx.SendStatus(201)
}

func (s *Server) AcumularPuntoHandler(ctx *fiber.Ctx) error {
	var modelo *Modelos.RequestAcumularPuntos
	err := ctx.BodyParser(&modelo)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON("json enviado es invalido")
	}
	errs := modelo.Validar()
	if len(errs) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(Modelos.Response{
			Errors: errs,
		})
	}
	err = s.servicioPuntos.ServicioAcumularPuntos(modelo)
	if err != nil {
		fmt.Println("Error en servicio acomulador puntos", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(Modelos.Response{
			Errors: []string{
				err.Error(),
			},
		})
	}
	return ctx.SendStatus(201)
}
