package cmd

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/x-mod/errors"
)

//default root command
var rootCmd *Command
var glogParse bool

func _program() string {
	return filepath.Base(os.Args[0])
}

//Add for default root command
func Add(opts ...CommandOpt) *Command {
	c := newCommand(opts...)
	c.build()
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

//GLOG support
func GLOG() {
	glogParse = true
}

//Execute for default root command
func Execute() {
	//add flag commandLine to support glog
	if glogParse {
		pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
		flag.Parse()
	}
	exit(rootCmd.Execute())
}

func exit(err error) {
	if err != nil {
		os.Exit(int(errors.ValueFrom(err)))
	}
}

func init() {
	rootCmd = newCommand(Name(_program()))
	rootCmd.TraverseChildren = true
	glogParse = false
}
