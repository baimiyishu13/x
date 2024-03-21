package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var mysqlExporterCmd = &cobra.Command{
	Use:   "mysqlExporter [user] [password]",
	Short: "Install MySQL Exporter",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		user := args[0]
		password := args[1]
		fmt.Println("Installing mysql exporter...")
		installMysqlExporter(user, password)
	},
}

func installMysqlExporter(user, password string) {
	cmd := exec.Command("bash", "-c", "mkdir -p /etc/mysqlexporter && echo -e \"[client]\nuser="+user+"\npassword="+password+"\" > /etc/mysqlexporter/.my.cnf")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to create .my.cnf file: %v\n", err)
		return
	}

	cmd = exec.Command("bash", "-c", "cp ./bin/mysqld_exporter  /usr/local/bin/")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Failed to decompress and move nmysqld_exporter : %v\n", err)
		return
	}
	service := `
[Unit]
Description=mysql Exporter
After=network.target

[Service]
Type=simple
ExecStart=/usr/local/bin/mysqld_exporter  --config.yaml.my-cnf=/etc/mysqlexporter/.my.cnf 

[Install]
WantedBy=multi-user.target
`
	cmd = exec.Command("bash", "-c", fmt.Sprintf("echo '%s' > /etc/systemd/system/mysql_exporter.service", service))
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Failed to create service file: %v\n", err)
		return
	}
	cmd = exec.Command("bash", "-c", "systemctl daemon-reload && systemctl start mysql_exporter.service && systemctl enable mysql_exporter.service")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Failed to start the service: %v\n", err)
		return
	}

	fmt.Println("🎉 Mysql exporter installed successfully.")

}

func init() {
	installCmd.AddCommand(mysqlExporterCmd)
}
