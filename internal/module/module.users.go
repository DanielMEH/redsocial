package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redsocial/internal/infrastructure/config"
	"github.com/redsocial/internal/infrastructure/constants"
	"github.com/redsocial/internal/infrastructure/types"
	"github.com/redsocial/internal/module/users/adapters/apis/handlers"
	"github.com/redsocial/internal/module/users/adapters/services/database"
	"github.com/redsocial/internal/module/users/application"
	"github.com/redsocial/internal/module/users/domains/ports"
	"go.uber.org/fx"
)

func configureRouterUser(register *handlers.RegisterAccountHandler, Hstore *types.HandlersStore, _ *config.AppSettings) {

	HandlerRouter := types.SliceHandlers{
		Prefix: "",
		Routes: []types.HandlerModule{
			{
				Route:   constants.API_ROUTER_STABLE + "/add_new_account",
				Method:  fiber.MethodPost,
				Handler: register.RunRegisterAccountHandler,
			},
		},
	}
	Hstore.Handlers = append(Hstore.Handlers, HandlerRouter)
}

func ModuleEmailsProvider() []fx.Option {
	return []fx.Option{

		// 1. Proveemos el Handler (el controlador)
		fx.Provide(handlers.NewRegisterAccountHandler),

		// 2. Dominios puertos
		fx.Provide(database.NewServicesDatabase,
			func(adapter *database.ServicesDatabaseAdapter) ports.PortsRepositoryUsers {
				return adapter
			}),

		// fx.Provide(usecases.NewUserUseCase),
		fx.Provide(application.NewRegisterAccountUseCase),

		// 3. Invocamos la configuración de rutas
		fx.Invoke(configureRouterUser),
	}
}
