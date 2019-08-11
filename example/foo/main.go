package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/x-mod/cmd"
)

func main() {
	svc := cmd.NewCommand(
		cmd.Name("svc"),
		cmd.SubCommand(
			cmd.NewCommand(
				cmd.Name("v1"),
				cmd.Main(V1),
			),
		),
	)
	cmd.Add(svc)
	cmd.Version("version string")
	cmd.Exit(cmd.Execute())
}

func V1(c *cmd.Command, args []string) {
	fmt.Println("V1 called", filepath.Base(os.Args[0]))
}
