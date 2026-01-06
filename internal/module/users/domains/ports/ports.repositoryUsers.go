package ports

import (
	"github.com/redsocial/internal/module/users/domains/command"
	"github.com/redsocial/internal/module/users/domains/entities/register"
)

type PortsRepositoryUsers interface {
	RegisterAccount(data command.EntityRegisterAccountCommand) (register.EntityRegisterAccountResponse, error)
}
