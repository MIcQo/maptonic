package cmd

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type OwnFormatter struct {
}

func (o OwnFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(entry.Message), nil
}

func init() {
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "debug mode")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose mode")
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
		// set up logging
		if verboseEnabled, err := cmd.Flags().GetBool("verbose"); err == nil && verboseEnabled {
			customFormatter := new(logrus.TextFormatter)
			customFormatter.TimestampFormat = "2006-01-02 15:04:05"
			customFormatter.FullTimestamp = true
			logrus.SetFormatter(customFormatter)
		} else {
			var customFormatter = new(OwnFormatter)
			logrus.SetFormatter(customFormatter)
		}

		if debugEnabled, err := cmd.Flags().GetBool("debug"); err == nil && debugEnabled {
			logrus.SetLevel(logrus.DebugLevel)
			logrus.Debug("Debug mode enabled")
		}

		// set up environment
		env := os.Getenv("APP_ENV")
		if "" == env {
			env = "development"
		}

		if err := loadEnv(env); err != nil {
			logrus.Error(err)
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

// loadEnv loads environment variables in specific order
// 1. .env.{env}.local
// 2. .env.{env}
// 3. .env.local (when is not `test` env)
// 4. .env
func loadEnv(env string) error {
	var firstLevelEnv = ".env." + env + ".local"
	logrus.WithFields(logrus.Fields{"env": firstLevelEnv}).Debug("Loading env..")
	var err = godotenv.Load(firstLevelEnv)
	if err != nil {
		logrus.Warn(err)
	}

	if "test" != env {
		var firstFallbackEnv = ".env.local"
		logrus.WithFields(logrus.Fields{"env": firstFallbackEnv}).Debug("Loading env..")
		err = godotenv.Load(firstFallbackEnv)
		if err != nil {
			logrus.Warn(err)
		}
	}

	var secondLevelEnv = ".env." + env
	logrus.WithFields(logrus.Fields{"env": secondLevelEnv}).Debug("Loading env..")
	err = godotenv.Load(secondLevelEnv)
	if err != nil {
		logrus.Warn(err)
	}

	err = godotenv.Load() // The Original .env
	return err
}
