package cmd

import (
	"super-indo-be/cmd/http"
	"super-indo-be/cmd/migration"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func Start() {
	rootCmd := &cobra.Command{}

	cmd := []*cobra.Command{
		{
			Use:   "serve-http",
			Short: "Run HTTP Server",
			Run: func(cmd *cobra.Command, args []string) {
				http.Start()
			},
		},
		{
			Use:   "db:migrate",
			Short: "Run DB migration related command",
			Run: func(cmd *cobra.Command, args []string) {
				migration.MigrateDatabase()
			},
		},
	}

	rootCmd.AddCommand(cmd...)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
