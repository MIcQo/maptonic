package cmd

import (
	"fmt"
	"runtime"

	"github.com/MIcQo/maptonic/config"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of MapTonic",
	Long:  `All software has versions. This is MapTonic's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(
			"MapTonic CLI version %s (built on %s - %s) \n",
			config.Version,
			runtime.Version(),
			runtime.GOARCH,
		)
	},
}

func init() {
	// Add versionCmd to the rootCmd
	rootCmd.AddCommand(versionCmd)
}
