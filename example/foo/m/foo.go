package m

import "github.com/x-mod/glog"

func Foo(i int) {
	glog.Info("foo info .... helo", i)
	glog.Warning("foo warning ... test", i)
	glog.Error("foo error .... loging", i)
}
