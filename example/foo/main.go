package main

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/spf13/viper"
	"github.com/x-mod/cmd"
)

func main() {
	cmd.Add(
		cmd.Path("/foo"),
		cmd.Main(V1),
	).PersistentFlags().StringP("parameter", "p", "test", "flags usage")
	cmd.Version("version string")
	cmd.Execute()
}

func V1(c *cmd.Command, args []string) error {
	defer glog.Flush()
	glog.Info("vlog info .... helo")
	glog.Warning("vlog warning ... test")
	glog.Error("vlog error .... loging")
	fmt.Println("V1 called, parameter:", viper.GetString("parameter"))
	glog.Info("ending vlog")
	return nil
}
