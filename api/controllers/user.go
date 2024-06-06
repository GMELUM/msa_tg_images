package controllers

import (
	"fmt"

	"github.com/gmelum/msa_tg_images/utils/avatar"
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

		identifier = fmt.Sprintf("user_%v", identifier)

		image, success := cache.Store.GetImage(identifier)
		if !success {
			file, err := images.GetUserImage(identifier, *config.TOKEN)
			if err != nil {
				file = avatar.CreateUser(identifier)
			}
			cache.Store.SetImage(identifier, file)
			image = &file
		}

		if string(*image)[:4] == "<svg" {
			ctx.Set("Content-Type", "image/svg+xml")
		} else {
			ctx.Set("Content-Type", "image/jpeg")
		}

		return ctx.Send(*image)

	})
}
