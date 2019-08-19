package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/x-mod/cmd"
)

func main() {
	cmd.Add(
		cmd.Name("svc"),
		cmd.Short("service short name"),
		cmd.SubCommand(
			cmd.NewCommand(
				cmd.Name("v1"),
				cmd.Main(V1),
			),
		),
	)
	cmd.Version("version string")
	cmd.Exit(cmd.Execute())
}

func V1(c *cmd.Command, args []string) {
	fmt.Println("V1 called", filepath.Base(os.Args[0]))
}
