package db

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
)

func NewConnection(host, port, user, dbname, password string) *sql.DB {
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
			host,
			port,
			user,
			dbname,
			password,
		),
	)
	if err != nil {
		logrus.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		logrus.Fatal(pingErr)
	}

	logrus.Info("Connected to database")
	return db
}
