package main

import (
	"github.com/gmelum/msa_tg_images/api"
	"github.com/gmelum/msa_tg_images/utils/cache"
	"github.com/gmelum/msa_tg_images/utils/config"

	"flag"

	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {

	flag.Parse()

	cache.New(cache.Option{
		Size:  *config.SIZE,
		Delay: *config.DELAY,
	})

	app := fiber.New(fiber.Config{
		AppName:                   "msa_cdn",
		DisableDefaultContentType: true,
	})

	api.NewRouter(app)

	panic(app.Listen(fmt.Sprintf(
		"%v:%v",
		"0.0.0.0",
		"18300",
	)))

}
