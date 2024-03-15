package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

/*
curl -L http://localhost:2379/metrics
*/
var etcdExporterCmd = &cobra.Command{
	Use:   "etcdExporter",
	Short: "Install etcd exporter",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("😘 write --> vmagent.yaml")
		installEtcdExporter()
	},
}

func installEtcdExporter() {
	conf := `
  - job_name: "etcd"
    scheme: https
    tls_config:
      insecure_skip_verify: true
      cert_file: /etc/ssl/etcd/ssl/node-dev-master.pem
      key_file: /etc/ssl/etcd/ssl/node-dev-master-key.pem
      ca_file: /etc/ssl/etcd/ssl/ca.pem
    static_configs:
      - targets: ['etcd-ip1:2379','etcd-ip2:2379','etcd-ip3:2379']`
	fmt.Println(conf)
}

func init() {
	installCmd.AddCommand(etcdExporterCmd)
}
