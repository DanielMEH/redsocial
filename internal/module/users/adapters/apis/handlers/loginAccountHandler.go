package handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/redsocial/internal/infrastructure/config"
	"github.com/redsocial/internal/module/users/application"
	"github.com/redsocial/internal/module/users/domains/command"
)

type LoginAccountHandler struct {
	useCase application.LoginAccountUseCase
}

func NewLoginAccountHandler(useCases *application.LoginAccountUseCase) *LoginAccountHandler {

	return &LoginAccountHandler{useCase: *useCases}
}

func (uc *LoginAccountHandler) RunLoginrAccountHandler(c *fiber.Ctx) error {

	var data command.EntityLoginAccountCommand

	if err := c.BodyParser(&data); err != nil {
		slog.Error(err.Error())
		return config.NewUnprStatusUnprocessableEntity(err)
	}

	responseUc, errorUsc := uc.useCase.ExecuteLoginAccountUseCase(data)

	if errorUsc != nil {
		return errorUsc
	}

	return config.ResponseOk(c, responseUc, "Ok")
}
