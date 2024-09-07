package cmd

import (
	"github.com/MIcQo/maptonic/internal/osmdata"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// importCmd represents the osmdata command
var importCmd = &cobra.Command{
	Use:   "import {osm_file}",
	Short: "Import OSM data using osm2pgsql",
	Long:  `Imports OpenStreetMap data into a PostgreSQL/PostGIS database using osm2pgsql.`,
	Args:  cobra.RangeArgs(0, 1), // Expect exactly one argument (the OSM file)
	Run: func(cmd *cobra.Command, args []string) {
		var osmFile = "./data/data.osm.pbf"
		if len(args) == 1 {
			osmFile = args[0]
		}

		dbHost, _ := cmd.Flags().GetString("host")
		dbPort, _ := cmd.Flags().GetString("port")
		dbUser, _ := cmd.Flags().GetString("user")
		dbName, _ := cmd.Flags().GetString("dbname")
		dbPassword, _ := cmd.Flags().GetString("password")
		update, _ := cmd.Flags().GetBool("update")

		if dbPassword == "" {
			logrus.Fatal("Error: Database password is required")
		}

		err := osmdata.Import(
			dbName,
			dbUser,
			dbHost,
			dbPort,
			dbPassword,
			osmFile,
			update,
		)
		if err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	// Add the import command to the root command
	rootCmd.AddCommand(importCmd)

	// Define flags for the import command
	importCmd.Flags().StringP("host", "H", "localhost", "Database host")
	importCmd.Flags().StringP("port", "P", "5432", "Database port")
	importCmd.Flags().StringP("user", "U", "", "Database user (required)")
	importCmd.Flags().StringP("dbname", "s", "", "Database (schema) name (required)")
	importCmd.Flags().StringP("password", "p", "", "Database password (required)")
	importCmd.Flags().BoolP("update", "u", false, "Update dataset")

	// Mark flags as required
	importCmd.MarkFlagRequired("user")
	importCmd.MarkFlagRequired("dbname")
	importCmd.MarkFlagRequired("password")
}
