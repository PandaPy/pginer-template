package main

import (
	"fmt"
	"os"

	"github.com/PandaPy/pginer/template/initialize/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println("命令执行失败:", err)
		os.Exit(1)
	}
}
