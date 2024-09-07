package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Dbname   string
	Password string
}

var db *bun.DB

func getDsn(host, port, user, dbname, password string) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		host,
		port,
		user,
		dbname,
		password,
	)
}

func NewConnection(config ...*Config) *bun.DB {
	if db != nil {
		return db
	}

	var host, port, user, dbname, password string
	if len(config) >= 1 {
		host = config[0].Host
		port = config[0].Port
		user = config[0].User
		dbname = config[0].Dbname
		password = config[0].Password
	} else {
		host = os.Getenv("DB_HOST")
		port = os.Getenv("DB_PORT")
		user = os.Getenv("DB_USER")
		dbname = os.Getenv("DB_NAME")
		password = os.Getenv("DB_PASSWORD")
	}

	sqldb := sql.OpenDB(
		pgdriver.NewConnector(pgdriver.WithDSN(getDsn(host, port, user, dbname, password))),
	)
	db = bun.NewDB(sqldb, pgdialect.New())

	pingErr := db.Ping()
	if pingErr != nil {
		logrus.Fatal(pingErr)
	}

	logrus.WithFields(logrus.Fields{
		"driver": db.String(),
	}).Debug("Connected to database")

	return db
}
