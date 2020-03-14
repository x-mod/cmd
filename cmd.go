package cmd

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

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

//Execute for default root command
func Execute() {
	//add flag commandLine to support glog
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	Exit(rootCmd.Execute())
}

//Exit with error code
func Exit(err error) {
	if err != nil {
		fmt.Println("failed: ", err)
		os.Exit(int(errors.ValueFrom(err)))
	}
}

func init() {
	rootCmd = newCommand(Name(_program()))
	rootCmd.TraverseChildren = true
}
