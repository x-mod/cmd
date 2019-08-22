package main

import (
	"github.com/x-mod/build"
	"github.com/x-mod/cmd"
	_ "github.com/x-mod/cmd/cmd/add"
	_ "github.com/x-mod/cmd/cmd/init"
)

func main() {
	cmd.Version(build.String())
	cmd.Execute()
}
