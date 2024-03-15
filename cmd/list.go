package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var components = []string{
	"🚀 vmagent",
	"🤕 vmalert",
	"🤕 victoriaMetrics",
	"🤕 grafnan",
	"🤕 prometheus",
	"🤕 alertmanager",
	"🚀 nodeExporter",
	"🚀 etcdExporter",
	"🚀 mysqlExporter",
	"🤕 mongoExporter",
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the installable components",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Installable components:")
		for _, component := range components {
			fmt.Printf("\t%s\n", component)
		}
	},
}

func init() {
	installCmd.AddCommand(listCmd)
}
