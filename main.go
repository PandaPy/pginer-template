package main

import (
	"os"

	"github.com/PandaPy/pginer/template/initialize"
)

func main() {
	if len(os.Args) > 1 {
		initialize.Cmd()
		return
	}
	initialize.RunWindowsServer()
}
