package main

import (
	"fmt"

	"github.com/x-mod/cmd"
)

func main() {
	cmd.RootMain(RootMain)
	cmd.Version("specified version")
	cmd.Exit(cmd.Execute())
}

func RootMain(c *cmd.Command, args []string) {
	fmt.Println("my root command running ...")
}
