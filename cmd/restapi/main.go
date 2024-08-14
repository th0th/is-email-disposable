package main

import (
	"github.com/go-errors/errors"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog/log"
	"github.com/th0th/is-email-disposable/pkg/isemaildisposable"
	"github.com/th0th/is-email-disposable/pkg/restapi/handler"
	"github.com/th0th/is-email-disposable/pkg/service/domain"
)

func main() {
	isemaildisposable.LogSetDefaults()

	domainService, err := domain.New()
	if err != nil {
		log.Panic().Stack().Err(errors.Wrap(err, 0)).Msg("email initialization failed")
	}

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Use(recover.New(recover.Config{EnableStackTrace: true, StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
		log.Error().Stack().Err(errors.Wrap(e.(error), 0)).Msg("fiber panic")
	}}))
	app.Use(fiberzerolog.New())

	app.Get("/", handler.Index(domainService))

	err = app.Listen("0.0.0.0:81")
	if err != nil {
		log.Panic().Stack().Err(errors.Wrap(err, 0)).Msg("fiber listen failed")
	}
}
