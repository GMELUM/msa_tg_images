package controllers

import (
	"fmt"

	"github.com/gmelum/msa_tg_images/utils/avatar"
	"github.com/gmelum/msa_tg_images/utils/cache"
	"github.com/gmelum/msa_tg_images/utils/config"
	"github.com/gmelum/msa_tg_images/utils/images"

	"github.com/gofiber/fiber/v2"
)

func Chat(path string, app fiber.Router) {
	app.Get(path, func(ctx *fiber.Ctx) error {

		identifier := ctx.Params("identifier")
		if len(identifier) == 0 {
			return fiber.NewError(fiber.StatusNotFound, "")
		}

		key := fmt.Sprintf("chat_%v", identifier)

		image, success := cache.Store.GetImage(key)
		if !success {
			file, err := images.GetChatImage(identifier, *config.TOKEN)
			if err != nil {
				file = avatar.CreateChat(key)
			}
			cache.Store.SetImage(key, file)
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
