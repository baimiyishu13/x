package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// etcdExporterCmd represents the etcdExporter command
var etcdExporterCmd = &cobra.Command{
	Use:   "etcdExporter",
	Short: "Install etcd exporter",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("etcdExporter called")
	},
}

func init() {
	installCmd.AddCommand(etcdExporterCmd)
}
