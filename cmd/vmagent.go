package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var vmagentCmd = &cobra.Command{
	Use:   "vmagent [dataDir] [vmip]",
	Short: "Install vmagent",
	Args:  cobra.ExactArgs(2), // Expect exactly 2 arguments
	Run: func(cmd *cobra.Command, args []string) {
		dataDir := args[0]
		vmip := args[1]
		installVMAgent(dataDir, vmip)
	},
}

func installVMAgent(dataDir string, vmip string) {
	fmt.Println("Installing vmagent...")
	// Step 1: Create data directory and decompress to it
	cmd := exec.Command("bash", "-c", fmt.Sprintf("mkdir -p %s && cp ./bin/vmagent-prod %s", dataDir, dataDir))
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to create data directory and decompress vmagent: %v\n", err)
		return
	}

	config := `global:
  external_labels:
    # 环境标签
    #datacenter: ceshi 
scrape_configs:
  #- job_name: node-exporter
  #  static_configs:
  #  - targets: ['1.1.1.1:9101','1.1.1.2:9101'] # 修改为监控的node节点IP
  #- job_name: mysql-exporter 
  #- job_name: etcd-exporter
`

	cmd = exec.Command("bash", "-c", fmt.Sprintf("echo '%s' > %s/vmagent.yml", config, dataDir))
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Failed to create configuration file: %v\n", err)
		return
	}

	// Step 3: Create service file
	service := fmt.Sprintf(`[Unit]
Description=vmagent
After=network.target

[Service]
Type=simple
ExecStart=%s/vmagent-prod -promscrape.config.yaml=%s/vmagent.yml -remoteWrite.url=http://%s:8428/api/v1/write
Restart=always

[Install]
WantedBy=multi-user.target
`, dataDir, dataDir, vmip)
	cmd = exec.Command("bash", "-c", fmt.Sprintf("echo '%s' > /etc/systemd/system/vmagent.service", service))
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Failed to create service file: %v\n", err)
		return
	}

	// Step 4: Start the service
	cmd = exec.Command("bash", "-c", "systemctl daemon-reload && systemctl restart vmagent.service  && systemctl enable vmagent.service")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Failed to start the service: %v\n", err)
		return
	}

	fmt.Println("🎉 VMAgent installed successfully.")
}

func init() {
	installCmd.AddCommand(vmagentCmd)
}
