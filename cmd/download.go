package cmd

import (
	"github.com/MIcQo/maptonic/internal/osmdata"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download {region}",
	Short: "Download OSM data from Geofabrik",
	Long:  `Downloads OpenStreetMap data from Geofabrik for a specified region.`,
	Args:  cobra.ExactArgs(1), // Expect exactly two arguments (region and output file path)
	Run: func(cmd *cobra.Command, args []string) {
		region := args[0] + "-latest.osm.pbf"

		if err := osmdata.Download(region); err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	// Add the download command to the root command
	rootCmd.AddCommand(downloadCmd)

	// Define flags for the download command
	// No additional flags needed for this basic implementation
}
