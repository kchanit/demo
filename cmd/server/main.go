package main

import (
	"context"

	"cloud.google.com/go/profiler"
	"github.com/centraldigital/cfw-cms-bff/infrastructure"
	"github.com/centraldigital/cfw-cms-bff/internal/adapter"
	"github.com/centraldigital/cfw-cms-bff/internal/core/service"
	"github.com/centraldigital/cfw-cms-bff/internal/handler"
	"github.com/centraldigital/cfw-cms-bff/internal/repository"
	"github.com/centraldigital/cfw-cms-bff/internal/repository/misc"
	"github.com/centraldigital/cfw-cms-bff/property"
	"github.com/centraldigital/cfw-core-lib/pkg/configuration/logger"
	"github.com/centraldigital/cfw-core-lib/pkg/configuration/server"
	"github.com/centraldigital/cfw-core-lib/pkg/configuration/tracer"
	"github.com/centraldigital/cfw-core-lib/pkg/configuration/validator"
	"github.com/centraldigital/cfw-core-lib/pkg/middleware"

	"github.com/centraldigital/cfw-cms-bff/internal/router"

	setupServer "github.com/centraldigital/cfw-core-lib/pkg/backoffice/setup-server"
)

func main() {
	ctx := server.ContextWithSignals(context.Background())

	// init logger
	log := logger.GetLogger()
	defer log.Sync()

	secretClient := infrastructure.NewSecretManagerClient(ctx)
	property.Init(ctx, secretClient)

	// Google Cloud Trace + OpenTelemetry init
	provider, err := tracer.InitCloudTraceProviderCtx(ctx)
	if err != nil {
		log.Fatalf("cannot init Cloud Trace provider: %v", err)
	}
	defer func() {
		if err := provider.Shutdown(ctx); err != nil {
			log.Errorf("cannot shutdown Cloud Trace provider: %v", err)
		}
	}()

	// setup profiler
	var profilerCfg profiler.Config
	if property.Get().Server.RunLocal {
		profilerCfg = profiler.Config{
			ProjectID: property.Get().Server.ProjectID,
			Service:   property.Get().Server.ServiceName,
		}
	}
	if err := profiler.Start(profilerCfg); err != nil {
		log.Panicf("unable to setup profiler: %v", err)
	}

	// init infrastructure
	pgx, scanapi := infrastructure.NewPostgres(ctx)

	// init adapter
	adpt := adapter.New()

	// init repository
	repo := repository.New(pgx, scanapi)

	// init miscrepo
	miscRepo, err := misc.New(ctx)

	// init service
	svc := service.New(adpt, repo, miscRepo)

	// create new gin engine
	engine := setupServer.InitServer()

	validator.AddDefaultValidators()

	// setup middleware
	middleware.NoLoggingMiddleware(engine)

	// init handler
	hdl := handler.New(svc)

	// setup router
	router.SetupRouter(engine, hdl)

	serverHost := property.Get().Server.Host
	serverPort := property.Get().Server.Port

	setupServer.StartServerWithCtx(ctx, engine, log, serverHost, serverPort)
}
