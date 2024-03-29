package m

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/x-mod/cmd"
	"github.com/x-mod/glog"
)

func Foo(i int) {
	glog.Info("foo info .... helo", i)
	glog.Warning("foo warning ... test", i)
	glog.Error("foo error .... loging", i)
}

func init() {
	c := cmd.Add(
		cmd.Path("/foo"),
		cmd.Main(V1),
	)
	c.Flags().StringP("parameter", "p", "test", "flags usage")
}

func V1(c *cmd.Command, args []string) error {
	glog.Open(
		glog.LogToStderr(true),
	)
	defer glog.Close()
	glog.Infoln("query:", viper.GetString("query"))
	glog.Infoln("parameter:", viper.GetString("parameter"))
	glog.MaxSize = 256
	return fmt.Errorf("error result")
}
