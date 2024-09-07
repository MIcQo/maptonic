/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/MIcQo/maptonic/cmd/data"

	"github.com/spf13/cobra"
)

// dataCmd represents the data command
var dataCmd = &cobra.Command{
	Use:   "data",
	Short: "Data related commands (import, download, update)",
	Long:  `Data related commands (import, download, update)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("data called")
	},
}

func init() {
	dataCmd.AddCommand(data.ImportCmd)
	dataCmd.AddCommand(data.DownloadCmd)
	rootCmd.AddCommand(dataCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dataCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dataCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
