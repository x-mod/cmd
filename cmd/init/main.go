package init

import (
	"github.com/x-mod/cmd"
)

func Main(c *cmd.Command, args []string) error {
	return nil
}

func init() {
	cmd.Add(
		cmd.Parent("/"),
		cmd.Name("init"),
		cmd.Short("init a command-line project"),
		cmd.Main(Main),
	)
}
