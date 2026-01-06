package handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/redsocial/internal/infrastructure/config"
	"github.com/redsocial/internal/module/users/application"
	"github.com/redsocial/internal/module/users/domains/command"
)

type RegisterAccountHandler struct {
	useCase application.RegisterAccountUseCase
}

func NewRegisterAccountHandler(useCases *application.RegisterAccountUseCase) *RegisterAccountHandler {

	return &RegisterAccountHandler{useCase: *useCases}
}

func (uc *RegisterAccountHandler) RunRegisterAccountHandler(c *fiber.Ctx) error {

	var data command.EntityRegisterAccountCommand

	if err := c.BodyParser(&data); err != nil {
		slog.Error(err.Error())
		return config.NewUnprStatusUnprocessableEntity(err)
	}

	responseUc, errorUsc := uc.useCase.ExecuteRegisterAccountUseCase(data)

	if errorUsc != nil {
		return errorUsc
	}

	return config.ResponseOk(c, responseUc, "Ok")
}
