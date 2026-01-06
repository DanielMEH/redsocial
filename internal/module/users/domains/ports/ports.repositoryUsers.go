package ports

import (
	"github.com/redsocial/internal/module/users/domains/command"
	"github.com/redsocial/internal/module/users/domains/entities/login"
	"github.com/redsocial/internal/module/users/domains/entities/profile"
	"github.com/redsocial/internal/module/users/domains/entities/register"
)

type PortsRepositoryUsers interface {
	RegisterAccount(data command.EntityRegisterAccountCommand) (register.EntityRegisterAccountResponse, error)
	LoginAccount(data command.EntityLoginAccountCommand) (login.EntityLoginAccountResponse, error)
	GetProfile(accountId string) (profile.EntityGetProfileAccountResponse, error)
}
