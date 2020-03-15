package cmd

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/golang/glog"

	"github.com/spf13/pflag"
	"github.com/x-mod/errors"
)

//default root command
var rootCmd *Command

//glog bool flag
var bGLOG bool

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
	bGLOG = true
}

//Execute for default root command
func Execute() {
	if bGLOG {
		pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	}
	exit(rootCmd.Execute())
}

func exit(err error) {
	if err != nil {
		glog.Error("executed: ", err)
	}
	os.Exit(int(errors.ValueFrom(err)))
}

func init() {
	rootCmd = newCommand(Name(_program()))
	rootCmd.TraverseChildren = true
	bGLOG = false
}
