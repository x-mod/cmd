package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func _default(c *cobra.Command, args []string) {
	fmt.Println("default called")
}

//CommandOpt command option definition
type CommandOpt func(*Command)

//Command struct
type Command struct {
	*cobra.Command
	childs []*Command
}

//NewCommand create a new command
func NewCommand(opts ...CommandOpt) *Command {
	cmd := &Command{
		Command: &cobra.Command{
			Run: _default,
		},
		childs: []*Command{},
	}
	for _, opt := range opts {
		opt(cmd)
	}
	return cmd
}

//Add subcommand
func (c *Command) Add(sub *Command) {
	for _, v := range c.childs {
		if v.Use == sub.Use {
			return
		}
	}
	c.childs = append(c.childs, sub)
	c.Command.AddCommand(sub.Command)
	c.Command.Run = nil
}

//PersistentFlags
func (c *Command) PersistentFlags() *pflag.FlagSet {
	return c.Command.PersistentFlags()
}

//Flags
func (c *Command) Flags() *pflag.FlagSet {
	return c.Command.Flags()
}

//Execute command
func (c *Command) Execute() error {
	if err := viper.BindPFlags(c.Flags()); err != nil {
		return err
	}
	if err := viper.BindPFlags(c.PersistentFlags()); err != nil {
		return err
	}
	return c.Command.Execute()
}

//Name of command
func Name(name string) CommandOpt {
	return func(cmd *Command) {
		cmd.Command.Use = name
	}
}

//Short Description of command
func Short(short string) CommandOpt {
	return func(cmd *Command) {
		cmd.Command.Short = short
	}
}

//Description of command
func Description(desc string) CommandOpt {
	return func(cmd *Command) {
		cmd.Command.Long = desc
	}
}

//MainFunc type
type MainFunc func(cmd *Command, args []string) error

//Main of command
func Main(main MainFunc) CommandOpt {
	return func(cmd *Command) {
		cmd.Command.Run = func(c *cobra.Command, args []string) {
			Exit(main(cmd, args))
		}
	}
}

//SubCommand of command
func SubCommand(sub *Command) CommandOpt {
	return func(cmd *Command) {
		cmd.Add(sub)
	}
}
