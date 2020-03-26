package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/x-mod/errors"
)

//default root command
var rootCmd *Command
var exitCode bool

func _program() string {
	return filepath.Base(os.Args[0])
}

//Add for default root command
func Add(opts ...CommandOpt) *Command {
	c := newCommand(opts...)
	c.build()
	if c.parent == nil {
		return rootCmd
	}
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

//ExitCode
func ExitCode(enable bool) {
	exitCode = enable
}

//Execute for default root command
func Execute() {
	if exitCode {
		exit(rootCmd.Execute())
	}
	rootCmd.Execute()
}

func exit(err error) {
	os.Exit(int(errors.ValueFrom(err)))
}

func init() {
	rootCmd = newCommand(Name(_program()))
	rootCmd.TraverseChildren = true
	exitCode = false
}
