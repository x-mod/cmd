package add

import (
	"github.com/x-mod/cmd"
)

func Main(c *cmd.Command, args []string) error {
	return nil
}

func init() {
	cmd.Add(
		cmd.Parent("/"),
		cmd.Name("add"),
		cmd.Short("add a subcommand"),
		cmd.Main(Main),
	)
}
