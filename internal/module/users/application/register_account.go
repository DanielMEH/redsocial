package application

import (
	"errors"

	"github.com/redsocial/internal/infrastructure/config"
	"github.com/redsocial/internal/module/users/domains/command"
	"github.com/redsocial/internal/module/users/domains/entities/register"
	"github.com/redsocial/internal/module/users/domains/ports"
)

type RegisterAccountUseCase struct {
	portsUsers ports.PortsRepositoryUsers
}

func NewRegisterAccountUseCase(portsUers ports.PortsRepositoryUsers) *RegisterAccountUseCase {

	return &RegisterAccountUseCase{portsUers}
}

func (uc *RegisterAccountUseCase) ExecuteRegisterAccountUseCase(data command.EntityRegisterAccountCommand) (register.EntityRegisterAccountResponse, error) {

	if _, err := register.NewEntityRegisterAccount(data); err != nil {

		return register.EntityRegisterAccountResponse{}, config.NewErrCodeEntitiesDataInvalid(errors.New(err.Error()))

	}

	responseUsc, err := uc.portsUsers.RegisterAccount(data)

	if err != nil {
		return register.EntityRegisterAccountResponse{}, err
	}

	return responseUsc, nil
}
