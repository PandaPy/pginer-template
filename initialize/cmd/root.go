// cmd/root.go

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pginer",
	Short: "Pginer CLI",
	Long:  `Pginer CLI 提供了启动服务器和管理用户的命令行工具`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Pginer CLI 已启动")
	},
}

// Execute 执行根命令
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(createSuperUserCmd)
}
