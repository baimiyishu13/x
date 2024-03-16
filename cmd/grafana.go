package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// grafanaCmd represents the grafana command
var grafanaCmd = &cobra.Command{
	Use:   "grafana",
	Short: "Install grafana",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("😘 CMD:")
		installGrafana()
	},
}

func installGrafana() {
	cmd := `
wget https://dl.grafana.com/enterprise/release/grafana-enterprise-10.3.3-1.x86_64.rpm
yum install -y ./grafana-enterprise-10.3.3-1.x86_64.rpm
systemctl start grafana-server.service
systemctl enable grafana-server.service`
	fmt.Println(cmd)
}

func init() {
	installCmd.AddCommand(grafanaCmd)
}
