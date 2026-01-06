package main

import (
	"github.com/joho/godotenv"
	"github.com/redsocial/internal/infrastructure/server"
	modules "github.com/redsocial/internal/module"
)

func main() {

	godotenv.Overload()

	app := server.ProviderServerStorage{}
	app.Init()
	app.AddModule(modules.ModuleEmailsProvider())
	app.Up()

}
