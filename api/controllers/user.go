package controllers

import (
	"github.com/gmelum/msa_tg_images/utils/cache"
	"github.com/gmelum/msa_tg_images/utils/config"
	"github.com/gmelum/msa_tg_images/utils/images"

	"github.com/gofiber/fiber/v2"
)

func User(path string, app fiber.Router) {
	app.Get(path, func(ctx *fiber.Ctx) error {

		identifier := ctx.Params("identifier")
		if len(identifier) == 0 {
			return fiber.NewError(fiber.StatusNotFound, "")
		}

		image, success := cache.Store.GetImage(identifier)
		if !success {
			file, err := images.GetUserImage(identifier, *config.TOKEN)
			if err != nil {
				return fiber.NewError(fiber.StatusNotFound, "")
			}
			cache.Store.SetImage(identifier, file)
			image = &file
		}

		ctx.Set("Content-Type", "image/jpeg")
		return ctx.Send(*image)

	})
}
