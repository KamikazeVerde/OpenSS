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

// modsCmd represents the mods command
var modsCmd = &cobra.Command{
	Use:   "mods",
	Short: "Scans for mods in .minecraft",
	Run: func(cmd *cobra.Command, args []string) {
		if strings.Contains(utils.CommandOutput("dir", utils.MinecraftFolder), "mods") {
			fmt.Println("Minecraft Fabric/Forge Mods:")
			utils.Command("dir", utils.MinecraftFolder+"/mods/")
		} else {
			fmt.Println("Mods folder not present")
		}
	},
}

func init() {
	rootCmd.AddCommand(modsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// modsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// modsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
