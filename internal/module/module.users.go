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

func configureRouterUser(
	register *handlers.RegisterAccountHandler,
	login *handlers.LoginAccountHandler,
	profile *handlers.GetProfileAccountHandler,
	Hstore *types.HandlersStore, _ *config.AppSettings) {

	HandlerRouter := types.SliceHandlers{
		Prefix: "",
		Routes: []types.HandlerModule{
			{
				Route:   constants.API_ROUTER_STABLE + "/add_new_account",
				Method:  fiber.MethodPost,
				Handler: register.RunRegisterAccountHandler,
			},
			{
				Route:   constants.API_ROUTER_STABLE + "/login",
				Method:  fiber.MethodPost,
				Handler: login.RunLoginrAccountHandler,
			},
			{
				Route:   constants.API_ROUTER_STABLE + "/profile",
				Method:  fiber.MethodGet,
				Handler: profile.RunGetProfileAccountHandler,
			},
		},
	}
	Hstore.Handlers = append(Hstore.Handlers, HandlerRouter)
}

func ModuleEmailsProvider() []fx.Option {
	return []fx.Option{

		// 1. Proveemos el Handler (el controlador)
		fx.Provide(handlers.NewRegisterAccountHandler),
		fx.Provide(handlers.NewGetProfileAccountHandler),
		fx.Provide(handlers.NewLoginAccountHandler),

		// 2. Dominios puertos
		fx.Provide(database.NewServicesDatabase,
			func(adapter *database.ServicesDatabaseAdapter) ports.PortsRepositoryUsers {
				return adapter
			}),

		// fx.Provide(usecases.NewUserUseCase),
		fx.Provide(application.NewRegisterAccountUseCase),
		fx.Provide(application.NewLoginAccountUseCase),
		fx.Provide(application.NewGetProfileAccountUseCase),

		// 3. Invocamos la configuración de rutas
		fx.Invoke(configureRouterUser),
	}
}
