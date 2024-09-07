package api

import (
	"fmt"
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

		logrus.WithFields(logrus.Fields{
			"lat": lat,
			"lon": lon,
		}).Debug("Reverse geocoding...")

		// TODO: Add reverse geocoding functionality
		fmt.Println(lat, lon)
	},
}
