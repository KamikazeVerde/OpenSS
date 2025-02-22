/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/kamikazeverde/OpenSS/utils"
	"github.com/spf13/cobra"
)

// lastsCmd represents the lasts command
var lastsCmd = &cobra.Command{
	Use:   "lasts",
	Short: "Ask for input and looks trough logs",
	Run: func(cmd *cobra.Command, args []string) {
		var input string
		if !utils.CheckPackage("lsof") {
			utils.InstallPackage("lsof")
		}
		fmt.Print("Input data: ")
		fmt.Scan(&input)
		if strings.Contains(utils.CommandOutput("lsof"), input) {
			fmt.Println("Input is present in logs")
		} else {
			fmt.Println("Input is not present in logs")
		}
	},
}

func init() {
	rootCmd.AddCommand(lastsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lastsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lastsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
