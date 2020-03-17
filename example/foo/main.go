package main

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/x-mod/cmd"
	"github.com/x-mod/cmd/example/foo/m"
	"github.com/x-mod/glog"
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
	glog.Open(
		glog.LogToStderr(true),
	)
	defer glog.Close()
	glog.Infoln("parameter:", viper.GetString("parameter"))
	glog.MaxSize = 256
	glog.Info("begin ... vlog")
	for i := 0; i < 16; i++ {
		glog.Info("vlog info .... xxxxx", i)
		glog.V(1).Info("vlog V(1) info ... pppppppp")
		glog.V(2).Info("vlog V(2) info ... good")
		glog.Warning("vlog warning ... yyyy", i)
		glog.Error("vlog error .... loging", i)
		m.Foo(i)
		fmt.Println("vlog looping: ", i)
	}
	return fmt.Errorf("error result")
}
