package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "debug mode")
}

// Root command definition
var rootCmd = &cobra.Command{
	Use:   "maptonic",
	Short: "MapTonic is a CLI tool for reverse geocoding and map tiles.",
	Long: `MapTonic CLI provides tools to work with reverse geocoding
and serve vector and raster map tiles.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Logic to handle the command, can be extended for additional features
		logrus.Info("Welcome to MapTonic!")
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		customFormatter := new(logrus.TextFormatter)
		customFormatter.TimestampFormat = "2006-01-02 15:04:05"
		customFormatter.FullTimestamp = true
		logrus.SetFormatter(customFormatter)

		if debugEnabled, err := cmd.Flags().GetBool("debug"); err == nil && debugEnabled {
			logrus.SetLevel(logrus.DebugLevel)
			logrus.Debug("Debug mode enabled")
		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
