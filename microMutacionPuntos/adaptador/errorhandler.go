package adaptador

import (
	"github.com/gofiber/fiber/v2"
)

var ErrorHandler = func(ctx *fiber.Ctx, err error) error {
	if err != nil {
		return err
	}
	return nil
}
