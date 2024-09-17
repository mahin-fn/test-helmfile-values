package cmd

import (
	"log"

	"github.com/mokhlesurr031/sre-project/backend/internal/connection"
	"github.com/mokhlesurr031/sre-project/backend/internal/migrations"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Starting Server...",
	Long:  `Starting Server...`,
	Run:   migrateDatabase,
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func migrateDatabase(cmd *cobra.Command, args []string) {
	// Ensure that the database connection is initialized
	if err := connection.ConnectDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	defer connection.CloseDB()

	// Run migrations
	if err := migrations.Migrate(); err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}

	log.Println("Database migration completed successfully")
}
