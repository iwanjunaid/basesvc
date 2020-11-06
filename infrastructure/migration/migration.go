package migration

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

type Migration struct {
	db *sql.DB
}

func (mig *Migration) RunMigration() {
	m, err := migrate.New("test", "Test")
	if err != nil {
		panic(err)
	}
	m.Steps()
	m, _ = migrate.New(
		"github://mattes:personal-access-token@mattes/migrate_test",
		"postgres://localhost:5432/database?sslmode=enable")
	m.Steps(2)

}
