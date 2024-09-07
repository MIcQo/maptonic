package cmd

import (
	"fmt"

	"github.com/MIcQo/maptonic/cmd/api"

	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "API related commands without starting HTTP server",
	Long:  `Here you can find commands related to the API. You can access standard API without running HTTP server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("api called")
	},
}

func init() {
	apiCmd.AddCommand(api.ServeCmd)
	apiCmd.AddCommand(api.ReverseCmd)
	apiCmd.AddCommand(api.SearchCmd)
	rootCmd.AddCommand(apiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
