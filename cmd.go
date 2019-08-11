package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/x-mod/errors"
)

//default root command
var rootCmd *Command

func _program() string {
	return filepath.Base(os.Args[0])
}

//Add for default root command
func Add(c *Command) {
	rootCmd.Add(c)
}

//Version root command
func Version(v string) {
	rootCmd.Version = v
}

//Main root command
func RootMain(rootFn MainFunc) {
	rootCmd.Run = func(c *cobra.Command, args []string) {
		rootFn(rootCmd, args)
	}
}

//Execute for default root command
func Execute() error {
	return rootCmd.Execute()
}

//Exit with error code
func Exit(err error) {
	if err != nil {
		os.Exit(int(errors.ValueFrom(err)))
	}
}

func init() {
	rootCmd = NewCommand(Name(_program()))
	rootCmd.TraverseChildren = true
}
