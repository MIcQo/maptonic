package api

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// SearchCmd represents the address search command
var SearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Geocode an address to get its latitude and longitude",
	Long:  `Takes an address as input and returns the corresponding latitude and longitude using geocoding.`,
	Args:  cobra.ExactArgs(1), // Expect exactly one argument (the address)
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Fatal("Not implemented yet")
	},
}
