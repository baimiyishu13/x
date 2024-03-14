package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var nodeExporterCmd = &cobra.Command{
	Use:   "nodeExporter",
	Short: "Install nodeExporter,default port 9101",

	Run: func(cmd *cobra.Command, args []string) {
		installNodeExporter()
	},
}

func installNodeExporter() {
	fmt.Println("Installing node exporter...")

	// Step 2: Decompress node_exporter and move it
	cmd := exec.Command("tar", "-zxvf", "node_exporter-1.7.0.linux-amd64.tar.gz", "-C", "/usr/local/bin/", "--strip-components=1")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to decompress and move node_exporter: %v\n", err)
		return
	}

	// Step 3: Create service file
	service := `
[Unit]
Description=Node Exporter
After=network.target

[Service]
Type=simple
ExecStart=/usr/local/bin/node_exporter  --web.listen-address=:9101
Restart=always

[Install]
WantedBy=multi-user.target
`
	cmd = exec.Command("bash", "-c", fmt.Sprintf("echo '%s' > /etc/systemd/system/node_exporter.service", service))
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Failed to create service file: %v\n", err)
		return
	}

	// Step 4: Start the service
	cmd = exec.Command("bash", "-c", "systemctl daemon-reload && systemctl start node_exporter.service && systemctl enable node_exporter.service")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Failed to start the service: %v\n", err)
		return
	}

	fmt.Println("🎉 Node exporter installed successfully.")
}

func init() {
	installCmd.AddCommand(nodeExporterCmd)
}
