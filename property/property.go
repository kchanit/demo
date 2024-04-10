package property

import (
	"context"
	"log"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	configs "github.com/centraldigital/cfw-core-lib/pkg/configuration/backoffice-property"
	"github.com/centraldigital/cfw-core-lib/pkg/util/infrastructureutil"
	"github.com/kelseyhightower/envconfig"
)

func Init(ctx context.Context, secretCli *secretmanager.Client) {
	configs.InitServerProperties()
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("read env error : %s", err.Error())
	}

	cfg.Postgres.PostgresConfig.SetConnPassword("postgres")
}

func Get() config {
	return cfg
}

var cfg config

type config struct {
	Server   serviceProperties
	Gin      gin
	Postgres postgres
}

type serviceProperties struct {
	configs.ServerProperties
}

type gin struct {
	Mode string `envconfig:"GIN_MODE" default:"debug"`
}

type postgres struct {
	PostgresConfig infrastructureutil.PostgresConfig
}
