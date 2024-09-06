package cmd

import (
	"github.com/MIcQo/maptonic/internal/api"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the HTTP server for MapTonic",
	Long:  `Starts an HTTP server to handle requests related to reverse geocoding and map tiles.`,
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt("port")
		if port == 0 {
			port = 8080 // default port
		}

		var debugEnabled, _ = cmd.Flags().GetBool("debug")
		var err = api.NewServer(&api.Config{
			Debug: debugEnabled,
			Host:  "0.0.0.0",
			Port:  uint(port),
		})

		if err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	// Add the serve command to the root command
	rootCmd.AddCommand(serveCmd)

	// Define flags for the serve command
	serveCmd.Flags().UintP("port", "p", 8080, "Port to run the HTTP server on")
}
