package migrations

import (
	"github.com/mokhlesurr031/sre-project/backend/domain"
	"github.com/mokhlesurr031/sre-project/backend/internal/connection"
)

// Migrate runs the migrations
func Migrate() error {
	// Auto-migrate all tables
	err := connection.DefaultDB().AutoMigrate(&domain.Resource{})
	if err != nil {
		return err
	}
	return nil
}
