package osmdata

import (
	"database/sql"
	"github.com/sirupsen/logrus"
)

func prepareDatabase(conn *sql.DB) error {
	logrus.Debug("Creating postgis extension...")
	if _, err := conn.Exec("CREATE EXTENSION IF NOT EXISTS postgis;"); err != nil {
		return err
	}

	logrus.Debug("Creating hstore extension...")
	if _, err := conn.Exec("CREATE EXTENSION IF NOT EXISTS hstore;"); err != nil {
		return err
	}

	return nil
}
