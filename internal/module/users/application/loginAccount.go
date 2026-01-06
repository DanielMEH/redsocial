package application

import (
	"errors"

	"github.com/redsocial/internal/infrastructure/config"
	"github.com/redsocial/internal/module/users/domains/command"
	"github.com/redsocial/internal/module/users/domains/entities/login"
	"github.com/redsocial/internal/module/users/domains/ports"
)

type LoginAccountUseCase struct {
	portsUsers ports.PortsRepositoryUsers
}

func NewLoginAccountUseCase(portsUers ports.PortsRepositoryUsers) *LoginAccountUseCase {

	return &LoginAccountUseCase{portsUers}
}

func (uc *LoginAccountUseCase) ExecuteLoginAccountUseCase(data command.EntityLoginAccountCommand) (login.EntityLoginAccountResponse, error) {

	if _, err := login.NewEntityLoginterAccount(data); err != nil {

		return login.EntityLoginAccountResponse{}, config.NewErrCodeEntitiesDataInvalid(errors.New(err.Error()))

	}

	responseUsc, err := uc.portsUsers.LoginAccount(data)

	if err != nil {
		return login.EntityLoginAccountResponse{}, err
	}

	return responseUsc, nil
}
