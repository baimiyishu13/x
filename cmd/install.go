package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a component",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("😘 Please check the command help")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
