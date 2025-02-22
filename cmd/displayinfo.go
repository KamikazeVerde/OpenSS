package cmd

import (
	"fmt"

	"github.com/kamikazeverde/OpenSS/utils"
	"github.com/spf13/cobra"
)

// displayinfoCmd represents the displayinfo command
var displayinfoCmd = &cobra.Command{
	Use:   "displayinfo",
	Short: "Displays system's information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Uname output:")
		utils.Command("uname", "-a")
		fmt.Println("\nUSB Devices:")
		if !utils.CheckPackage("usbutils") {
			utils.InstallPackage("usbutils")
		} else {
			utils.Command("lsusb")
		}
		fmt.Println("\nDisk Info:")
		utils.Command("lsblk")
		fmt.Println("\nPCI Devices List:")
		if !utils.CheckPackage("pciutils") {
			utils.InstallPackage("pciutils")
		} else {
			utils.Command("lspci")
		}
		fmt.Println("\nBoot time:")
		utils.Command("who", "-b")

		fmt.Println("\nMinecraft Directory:")
		fmt.Println(utils.MinecraftFolder)
	},
}

func init() {
	rootCmd.AddCommand(displayinfoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// displayinfoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// displayinfoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
