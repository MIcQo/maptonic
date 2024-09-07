package api

import (
	"github.com/MIcQo/maptonic/internal/api"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ServeCmd represents the serve command
var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the HTTP server for MapTonic",
	Long:  `Starts an HTTP server to handle requests related to reverse geocoding and map tiles.`,
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetUint("port")
		if port == 0 {
			port = 8080 // default port
		}

		var debugEnabled, _ = cmd.Flags().GetBool("debug")
		var err = api.NewServer(&api.Config{
			Debug: debugEnabled,
			Host:  "0.0.0.0",
			Port:  port,
		})

		if err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	// Define flags for the serve command
	ServeCmd.Flags().UintP("port", "p", 8080, "Port to run the HTTP server on")
}
