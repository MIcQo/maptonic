package db

import (
	"database/sql"
	"fmt"
	"github.com/oiime/logrusbun"
	"github.com/uptrace/bun/extra/bundebug"
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
	Debug    bool
}

var db *bun.DB

func NewConnection(config ...*Config) *bun.DB {
	if db != nil {
		return db
	}

	var host = os.Getenv("DB_HOST")
	var port = os.Getenv("DB_PORT")
	var user = os.Getenv("DB_USER")
	var dbname = os.Getenv("DB_NAME")
	var password = os.Getenv("DB_PASSWORD")
	var debug = false

	if len(config) >= 1 {
		var c = config[0]
		if c.Host != "" {
			host = c.Host
		}

		if c.Port != "" {
			port = c.Port
		}

		if c.User != "" {
			user = c.User
		}

		if c.Dbname != "" {
			dbname = c.Dbname
		}

		if c.Password != "" {
			password = c.Password
		}

		debug = c.Debug
	}

	sqldb := sql.OpenDB(
		pgdriver.NewConnector(
			pgdriver.WithAddr(fmt.Sprintf("%s:%s", host, port)),
			pgdriver.WithUser(user),
			pgdriver.WithPassword(password),
			pgdriver.WithDatabase(dbname),
			pgdriver.WithTLSConfig(nil),
		),
	)

	db = bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(debug)))
	db.AddQueryHook(logrusbun.NewQueryHook(logrusbun.QueryHookOptions{Logger: logrus.StandardLogger()}))

	pingErr := db.Ping()
	if pingErr != nil {
		logrus.Fatal(pingErr)
	}

	logrus.WithFields(logrus.Fields{
		"driver": db.String(),
	}).Debug("Connected to database")

	return db
}
