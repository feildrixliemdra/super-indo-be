package migration

import (
	"super-indo-be/internal/bootstrap"
	"super-indo-be/pkg/dbmigration"
)

func MigrateDatabase() {
	cfg := bootstrap.NewConfig()

	dbmigration.DatabaseMigration(cfg)
}
