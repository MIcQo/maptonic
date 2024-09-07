package osmdata

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/MIcQo/maptonic/config"
	"github.com/sirupsen/logrus"
)

func getUriForRegion(region string) string {
	return fmt.Sprintf("https://download.geofabrik.de/%s", region)
}

// Download - Download OSM data from Geofabrik region should not contain "-latest" or file extension.
func Download(region string) error {
	var url = getUriForRegion(region)
	logrus.Infof("Downloading OSM data for region %s from %s...\n", region, url)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to download file: %s\n", resp.Status)
	}

	outFile, err := os.Create(config.OSMDataDir)
	if err != nil {
		return fmt.Errorf("Error creating file: %s\n", err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return fmt.Errorf("Error writing file: %s\n", err)
	}

	logrus.Infof("OSM data downloaded successfully to %s\n", config.OSMDataDir)
	return nil
}
