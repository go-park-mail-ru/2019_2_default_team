package db

import (
	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
	"kino_backend/logger"
)

func MakeMigrations(db *sqlx.DB) {
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	n, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Up)
	if err != nil {
		logger.Error(err)
	} else if n != 0 {
		logger.Infof("Applied %d migrations!", n)
	}
}
