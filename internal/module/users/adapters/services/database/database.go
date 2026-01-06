package database

import (
	"github.com/google/uuid"
	"github.com/redsocial/internal/infrastructure/config"
	"github.com/redsocial/internal/module/users/adapters/utils"
	"github.com/redsocial/internal/module/users/domains/command"
	"github.com/redsocial/internal/module/users/domains/entities/register"
)

type ServicesDatabaseAdapter struct {
}

func NewServicesDatabase() *ServicesDatabaseAdapter {
	return &ServicesDatabaseAdapter{}
}

func (sd *ServicesDatabaseAdapter) RegisterAccount(data command.EntityRegisterAccountCommand) (register.EntityRegisterAccountResponse, error) {

	uuidValue := uuid.New().String()

	value, errJwt := utils.GenerateJWT(uuidValue, uuidValue)

	if errJwt != nil {
		println("error")
		return register.EntityRegisterAccountResponse{}, config.NewInternalServerError(errJwt)
	}

	return register.EntityRegisterAccountResponse{
		Message: "ok",
		Details: struct {
			AccountToken     string "json:\"account_token\""
			AccountSessionId string "json:\"account_session_id\""
		}{
			AccountToken:     value,
			AccountSessionId: uuidValue,
		},
	}, nil
}
