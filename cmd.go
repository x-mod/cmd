package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/x-mod/errors"
)

//default root command
var rootCmd *Command

func _program() string {
	return filepath.Base(os.Args[0])
}

//Add for default root command
func Add(opts ...CommandOpt) *Command {
	c := NewCommand(opts...)
	rootCmd.Add(c)
	return c
}

//PersistentFlags
func PersistentFlags() *pflag.FlagSet {
	return rootCmd.PersistentFlags()
}

//Flags
func Flags() *pflag.FlagSet {
	return rootCmd.Flags()
}

//Version root command
func Version(v string) {
	rootCmd.Version = v
}

//Main root command
func RootMain(rootFn MainFunc) {
	rootCmd.Run = func(c *cobra.Command, args []string) {
		Exit(rootFn(rootCmd, args))
	}
}

//Execute for default root command
func Execute() error {
	return rootCmd.Execute()
}

//Exit with error code
func Exit(err error) {
	if err != nil {
		fmt.Println("failed: ", err)
		os.Exit(int(errors.ValueFrom(err)))
	}
}

func init() {
	rootCmd = NewCommand(Name(_program()))
	rootCmd.TraverseChildren = true
}
