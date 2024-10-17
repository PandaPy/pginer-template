// cmd/root.go

package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Pginer",
	Short: "Pginer CLI",
	Long:  `Pginer CLI 提供了启动服务器和管理用户的命令行工具`,
	Run: func(cmd *cobra.Command, args []string) {
		// 显示帮助信息
		_ = cmd.Help()
	},
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.SetHelpCommand(&cobra.Command{
		Hidden: true,
	})
}
