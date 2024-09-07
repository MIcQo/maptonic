package api

import (
	"github.com/MIcQo/maptonic/internal/api"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
)

// ReverseCmd represents the reverse geocode command
var ReverseCmd = &cobra.Command{
	Use:   "reverse {lat} {lon}",
	Short: "Reverse geocode a latitude and longitude",
	Long:  `Takes latitude and longitude as input and returns the nearest address using reverse geocoding.`,
	Args:  cobra.ExactArgs(2), // Expect exactly two arguments (latitude and longitude)
	Run: func(cmd *cobra.Command, args []string) {
		var lat, latErr = strconv.ParseFloat(args[0], 64)
		if latErr != nil {
			logrus.Fatal(latErr)
		}

		var lon, lonErr = strconv.ParseFloat(args[1], 64)
		if lonErr != nil {
			logrus.Fatal(lonErr)
		}

		var zoom, _ = cmd.Flags().GetUint("zoom")
		logrus.WithFields(logrus.Fields{
			"lat":  lat,
			"lon":  lon,
			"zoom": zoom,
		}).Debug("Reverse geocoding...")

		var result, err = api.ReverseGeocode(lat, lon, zoom)
		if err != nil {
			logrus.Fatal(err)
		}

		logrus.Infof("%+v", result)
	},
}

func init() {
	ReverseCmd.Flags().UintP("zoom", "z", 13, "Zoom level")
}
