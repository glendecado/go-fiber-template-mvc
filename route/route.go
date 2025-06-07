package route

import (
	"github.com/glendecado/template/controller"
	"github.com/gofiber/fiber/v2"
)

func Run(app *fiber.App) {

	//sample route
	app.Get("/", controller.Hello)
}
