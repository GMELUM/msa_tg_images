package api

import (
	"github.com/gmelum/msa_tg_images/api/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewRouter(app *fiber.App) {

	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))

	controllers.User("/user/:identifier", app)
	controllers.Chat("/chat/:identifier", app)

}
