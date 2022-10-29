package commands

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"strings"
)

func Migrate(log *zap.SugaredLogger, command string, connectionString string, migrationDir string) error {
	if command != "up" && command != "down" {
		return errors.New("command must be up or down")
	}
	if connectionString == "" {
		return errors.New("pass a valid mongodb connection string")
	}
	if migrationDir == "" || !strings.HasPrefix(migrationDir, "file://") {
		return errors.New("migration directory must start with file://, example file://./myfolder")
	}
	migration, err := migrate.New(migrationDir, connectionString)
	if err != nil {
		return err
	}

	fn := migration.Down
	if command == "up" {
		fn = migration.Up
	}

	err = fn()
	if err != nil {
		if err != migrate.ErrNoChange {
			return err
		}
		log.Warn(err)
	}

	return nil
}
