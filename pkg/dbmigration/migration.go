package dbmigration

import (
	"context"
	"flag"
	"fmt"
	"os"
	"super-indo-be/internal/config"

	"github.com/pressly/goose/v3"
	log "github.com/sirupsen/logrus"
)

var (
	flags   = flag.NewFlagSet("db:migrate", flag.ExitOnError)
	dir     = flags.String("dir", "database/migration", "directory with migration files")
	table   = flags.String("table", "db_migration", "migrations table name")
	help    = flags.Bool("guide", false, "print help")
	verbose = flags.Bool("verbose", false, "set verbose")
)

func DatabaseMigration(cfg *config.Config) {
	flags.Usage = usage
	flags.Parse(os.Args[2:])
	ctx := context.Background()

	goose.SetVerbose(*verbose)
	goose.SetTableName(*table)

	args := flags.Args()

	if len(args) == 0 || *help {
		flags.Usage()
		return
	}

	switch args[0] {
	case "create":
		if err := goose.RunContext(ctx, "create", nil, *dir, args[1:]...); err != nil {
			log.Fatalf("goose run: %v", err)
		}
		return
	case "fix":
		if err := goose.RunContext(ctx, "fix", nil, *dir); err != nil {
			log.Fatalf("goose run: %v", err)
		}
		return
	}

	if len(args) < 1 {
		flags.Usage()
		return
	}

	command := args[0]
	fmt.Println("db url ", cfg.Postgre.URL)
	db, err := goose.OpenDBWithDriver("postgres", cfg.Postgre.URL)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("db migrate: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}

	if err := goose.RunContext(ctx, command, db, *dir, arguments...); err != nil {
		log.Fatalf("db migrate run: %v", err)
	}
}

func usage() {
	fmt.Println(usageCommands)
}

var (
	usageCommands = `
  --dir string     directory with migration files (default "database/migration")
  --guide          print help
  --table string   migrations table name (default "db_migration")
  --verbose        enable verbose mode
  --version        print version

Commands:
    up                   Migrate the DB to the most recent version available
    up-by-one            Migrate the DB up by 1
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    reset                Roll back all migrations
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with the current timestamp
    fix                  Apply sequential ordering to migrations
`
)
