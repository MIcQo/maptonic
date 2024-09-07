package osmdata

import (
	"bytes"
	"fmt"
	"github.com/MIcQo/maptonic/internal/db"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
)

func Import(dbName, dbUser, dbHost, dbPort, dbPassword, osmFile string, update bool) error {
	var conn = db.NewConnection(dbHost, dbPort, dbUser, dbName, dbPassword)
	defer conn.Close()

	if err := prepareDatabase(conn); err != nil {
		return err
	}

	if err := importOsmFile(dbName, dbUser, dbHost, dbPort, dbPassword, osmFile, update); err != nil {
		return err
	}

	if err := importCountryGrid(dbUser, dbPassword, dbHost, dbPort, dbName); err != nil {
		return err
	}

	//if err := importPlacex(conn); err != nil {
	//	return err
	//}

	return nil
}

func importCountryGrid(user, password, host, port, dbName string) error {
	var sqlFilePath = "./migrations/country_grid.sql"
	cmd := exec.Command(
		"psql",
		"-U",
		user,
		"-p",
		port,
		"-h",
		host,
		"-d",
		dbName,
		"-a",
		"-f",
		sqlFilePath,
	)

	logrus.Debug(cmd.String())
	cmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", password))
	var out, stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Error executing query. Command Output: %+v\n: %+v, %v", out.String(), stderr.String(), err)
	}

	logrus.Debug("country_grid.sql executed successfully")
	return nil
}

//
//func importPlacex(db *sql.DB) error {
//	var f, err = os.Open("./migrations/placex.sql")
//	if err != nil {
//		return err
//	}
//	c, ioErr := io.ReadAll(f)
//	if ioErr != nil {
//		return ioErr
//	}
//
//	_, err = db.Exec(string(c))
//	if err != nil {
//		return err
//	}
//
//	logrus.Debug("placex.sql executed successfully")
//	return nil
//}
