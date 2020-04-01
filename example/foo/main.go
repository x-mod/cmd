package main

import (
	"github.com/x-mod/cmd"

	_ "github.com/x-mod/cmd/example/foo/m"
)

func main() {
	cmd.Add(
		cmd.Name("foo"),
	)
	cmd.PersistentFlags().StringP("query", "q", "dddd", "flags usage")
	cmd.Execute()
}
