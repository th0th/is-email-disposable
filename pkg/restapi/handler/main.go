package handler

import (
	"github.com/go-errors/errors"
	"github.com/gofiber/fiber/v2"

	"github.com/th0th/is-email-disposable/pkg/isemaildisposable"
)

type indexQueryParams struct {
	EmailOrDomain *string `query:"email"`
}

func Index(domainService isemaildisposable.DomainService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		queryParams := indexQueryParams{}
		err := c.QueryParser(&queryParams)
		if err != nil {
			return errors.Wrap(err, 0)
		}

		if queryParams.EmailOrDomain == nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"email": "This parameter is required."})
		}

		checkResult := domainService.Check(*queryParams.EmailOrDomain)

		err = c.JSON(checkResult)
		if err != nil {
			return errors.Wrap(err, 0)
		}

		return nil
	}
}
