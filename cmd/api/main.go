package main

import (
	"fmt"

	"github.com/radityacandra/sekolahku-smanssa-api/configs"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/server"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/server/types"
	"github.com/radityacandra/sekolahku-smanssa-api/pkg/database"
)

func main() {
	// load default config
	config, err := configs.LoadGlobalConfig()
	if err != nil {
		panic(fmt.Errorf("cannot load default config with error: %s", err))
	}

	// load db connection
	db, err := database.NewDatabase(config.PostgresUri)
	if err != nil {
		panic(fmt.Errorf("cannot load database with error: %s", err))
	}

	dep := types.NewDependency(db, config)

	// start web server
	server.InitServer(dep)
}
