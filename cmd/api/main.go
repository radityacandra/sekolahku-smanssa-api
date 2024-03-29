package main

import (
	"fmt"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/radityacandra/sekolahku-smanssa-api/configs"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/server"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/server/types"
	"github.com/radityacandra/sekolahku-smanssa-api/pkg/database"
	"github.com/radityacandra/sekolahku-smanssa-api/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	// load logger
	logger := logger.LoadLogger()

	// load default config
	config, err := configs.LoadGlobalConfig(logger)
	if err != nil {
		panic(fmt.Errorf("cannot load default config with error: %s", err))
	}

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("Sekolahku Smanssa API"),
		newrelic.ConfigLicense(config.NewRelicLisenceKey),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		logger.Error("Failed to initiate newrelic", zap.Error(err))
	}

	// load db connection
	db, err := database.NewDatabase(config.PostgresUri, logger)
	if err != nil {
		panic(fmt.Errorf("cannot load database with error: %s", err))
	}

	dep := types.NewDependency(db, config, logger, app)

	// start web server
	server.InitServer(dep)
}
