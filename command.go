package cmd

import (
	"flag"
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/golang/glog"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func _default(c *cobra.Command, args []string) {
	fmt.Println("default called with args:", args)
}

//CommandOpt command option definition
type CommandOpt func(*Command)

//Command struct
type Command struct {
	*cobra.Command
	dir    string
	name   string
	parent *Command
	childs map[string]*Command
}

//create a new command
func newCommand(opts ...CommandOpt) *Command {
	cmd := &Command{
		Command: &cobra.Command{
			Run: _default,
		},
		parent: nil,
		childs: make(map[string]*Command),
	}
	for _, opt := range opts {
		opt(cmd)
	}
	return cmd
}

//build relations
func (c *Command) build() {
	//replace rootCmd
	if c.dir == "" {
		if c.name != "" {
			rootCmd.name = c.name
			rootCmd.Command.Use = c.name
		}
		if c.Command.Short != "" {
			rootCmd.Command.Short = c.Command.Short
		}
		if c.Command.Long != "" {
			rootCmd.Command.Long = c.Command.Long
		}
		if len(rootCmd.childs) == 0 {
			rootCmd.Command.Run = c.Command.Run
		}
		return
	}
	parent := rootCmd
	if c.dir != "/" {
		dirs := strings.Split(strings.TrimPrefix(c.dir, "/"), "/")
		for _, d := range dirs {
			if ci, ok := parent.childs[d]; ok {
				parent = ci
			} else {
				nc := newCommand()
				nc.setName(d)
				nc.parent = parent
				nc.dir = path.Join(nc.parent.dir, nc.parent.name)
				parent.Command.Run = nil
				parent.AddCommand(nc.Command)
				parent.childs[d] = nc
				parent = nc
			}
		}
	}
	c.parent = parent
	c.parent.Command.Run = nil
	c.parent.AddCommand(c.Command)
	c.parent.childs[c.name] = c
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
	for _, child := range c.childs {
		if err := viper.BindPFlags(child.Flags()); err != nil {
			return err
		}
		if err := viper.BindPFlags(child.PersistentFlags()); err != nil {
			return err
		}
	}
	return c.Command.Execute()
}

//Path of command
func Path(p string) CommandOpt {
	return func(cmd *Command) {
		if len(p) > 0 {
			if p == "/" {
				cmd.dir = ""
				cmd.setName(_program())
			} else {
				dir, name := filepath.Split(p)
				cmd.dir = dir
				if dir != "/" {
					cmd.dir = strings.TrimSuffix(dir, "/")
				}
				cmd.setName(name)
			}
		}
	}
}

func (cmd *Command) setName(name string) {
	cmd.name = name
	cmd.Command.Use = name
	cmd.Command.Short = fmt.Sprintf("%s command", name)
}

//Parent of command
func Parent(parent string) CommandOpt {
	return func(cmd *Command) {
		cmd.dir = parent
	}
}

//Name of command
func Name(name string) CommandOpt {
	return func(cmd *Command) {
		cmd.name = name
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
			if bGLOG {
				defer glog.Flush()
				flag.Parse()
				exit(main(cmd, args))
			}
			exit(main(cmd, args))
		}
	}
}
