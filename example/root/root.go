package main

import (
	"fmt"

	"github.com/x-mod/cmd"
)

func main() {
	cmd.Add(
		cmd.Name("root"),
		cmd.Main(RootMain),
	)
	cmd.Execute()
}

func RootMain(c *cmd.Command, args []string) error {
	fmt.Println("my root command running ...")
	return nil
}
