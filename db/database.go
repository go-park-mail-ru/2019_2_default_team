package db

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver
	"kino_backend/logger"
)

var Db *sqlx.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "docker"
	dbname   = "postgres"
)

func InitDB(address, database string) *sqlx.DB {
	var err error

	//dsn := "localhost:5432/some-postgres?"
	//Db, err = sqlx.Open("postgres",
	//	"postgres://"+address+"/"+database+"?sslmode=disable")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	Db, err = sqlx.Open("postgres", psqlInfo)

	if err != nil {
		logger.Panic(err)
	}

	if err := Db.Ping(); err != nil {
		logger.Panic(err)
	}

	//Db, err := sqlx.Connect("postgres", "user=foo dbname=some-bar sslmode=disable")
	//if err != nil {
	//	log.Fatalln(err)
	//}

	logger.Infof("Successfully connected to %v, database %v", address, database)

	MakeMigrations(Db)

	return Db
}

func InitDBSQL(address, database string) *sql.DB {
	var err error

	//dsn := "localhost:5432/some-postgres?"
	//Db, err = sqlx.Open("postgres",
	//	"postgres://"+address+"/"+database+"?sslmode=disable")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	Database, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		logger.Panic(err)
	}

	if err := Db.Ping(); err != nil {
		logger.Panic(err)
	}

	//Db, err := sqlx.Connect("postgres", "user=foo dbname=some-bar sslmode=disable")
	//if err != nil {
	//	log.Fatalln(err)
	//}

	logger.Infof("Successfully connected to %v, database %v", address, database)

	MakeMigrations(Db)

	return Database
}
