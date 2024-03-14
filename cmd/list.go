/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var components = []string{
	"vmagent",
	"node-exporter",
	"etcd-exporter",
	"mysql-exporter",
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the installable components",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Installable components:")
		for _, component := range components {
			fmt.Printf("\t %s\n", component)
		}
	},
}

func init() {
	installCmd.AddCommand(listCmd)
}
