package main

import (
	"fmt"

	"github.com/x-mod/cmd"
)

func main() {
	cmd.Add(
		cmd.Parent("/service/foo/bar"),
		cmd.Name("v1"),
		cmd.Main(V1),
	).PersistentFlags().StringP("parameter", "p", "test", "flags usage")
	cmd.Version("version string")
	cmd.Execute()
}

func V1(c *cmd.Command, args []string) error {
	fmt.Println("V1 called")
	return nil
}
