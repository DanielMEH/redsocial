package handlers

import (
	"errors"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/redsocial/internal/infrastructure/config"
	"github.com/redsocial/internal/module/users/application"
)

type GetProfileAccountHandler struct {
	useCase application.GetProfileAccountUseCase
}

func NewGetProfileAccountHandler(useCases *application.GetProfileAccountUseCase) *GetProfileAccountHandler {

	return &GetProfileAccountHandler{useCase: *useCases}
}

func (uc *GetProfileAccountHandler) RunGetProfileAccountHandler(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	// 2. Validar que no venga vacío y tenga el prefijo "Bearer "
	if authHeader == "" || len(authHeader) < 8 || authHeader[:7] != "Bearer " {
		slog.Error("token no proporcionado o formato inválido no authorizado")
		return config.NewStatusUnauthorized(errors.New("token no proporcionado o formato inválido"))
	}
	tokenString := authHeader[7:]
	responseUc, errorUsc := uc.useCase.ExecuteGetProfileccountUseCase(tokenString)

	if errorUsc != nil {
		return errorUsc
	}

	return config.ResponseOk(c, responseUc, "Ok")
}
