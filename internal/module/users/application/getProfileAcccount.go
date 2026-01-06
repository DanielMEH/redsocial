package application

import (
	"github.com/redsocial/internal/infrastructure/config"
	"github.com/redsocial/internal/infrastructure/utils"
	"github.com/redsocial/internal/module/users/domains/entities/profile"
	"github.com/redsocial/internal/module/users/domains/ports"
)

type GetProfileAccountUseCase struct {
	portsUsers ports.PortsRepositoryUsers
}

func NewGetProfileAccountUseCase(portsUers ports.PortsRepositoryUsers) *GetProfileAccountUseCase {

	return &GetProfileAccountUseCase{portsUers}
}

func (uc *GetProfileAccountUseCase) ExecuteGetProfileccountUseCase(tokenString string) (profile.EntityGetProfileAccountResponse, error) {

	accountId, errToken := utils.ValidateJWT(tokenString)

	if errToken != nil {
		return profile.EntityGetProfileAccountResponse{}, config.NewStatusUnauthorized(errToken)
	}
	responseUsc, err := uc.portsUsers.GetProfile(accountId)

	if err != nil {
		return profile.EntityGetProfileAccountResponse{}, err
	}

	return responseUsc, nil
}
