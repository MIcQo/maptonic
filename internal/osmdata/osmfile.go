package osmdata

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/sirupsen/logrus"
)

func importOsmFile(dbName, dbUser, dbHost, dbPort, dbPassword, osmFile string, update bool) error {
	// Build osm2pgsql command
	cmdArgs := []string{
		"--database", dbName,
		"--user", dbUser,
		"--host", dbHost,
		"--port", dbPort,
		"--slim",
		"-C", "10000",
		"-W",
		"-l",
		"--hstore-all",
	}

	if update {
		cmdArgs = append(cmdArgs, "-a")
	} else {
		cmdArgs = append(cmdArgs, "-c")
	}

	cmdArgs = append(cmdArgs, osmFile)

	osm2pgsqlCmd := exec.Command("osm2pgsql", cmdArgs...)
	osm2pgsqlCmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", dbPassword))
	logrus.Debug(osm2pgsqlCmd.String())
	osm2pgsqlCmd.Stdout = os.Stdout
	osm2pgsqlCmd.Stderr = os.Stderr

	logrus.Infof("Importing OSM data from %s into database %s...\n", osmFile, dbName)

	if err := osm2pgsqlCmd.Run(); err != nil {
		return err
	}

	logrus.Info("OSM data import completed successfully.")
	return nil
}
